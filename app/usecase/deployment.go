package usecase

import (
	"fmt"
	"github.com/goal-web/collection"
	"github.com/goal-web/contracts"
	utils2 "github.com/goal-web/supports/utils"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type deploymentCommand func(deployment DeploymentDetail, server models.Server, script string) (string, error)

var tempRepoPath string
var deploymentsChan = make(map[int]chan DeploymentParam)

func init() {
	tempRepoPath = os.TempDir()
}

func deploymentChan(projectId int) chan DeploymentParam {
	if ch := deploymentsChan[projectId]; ch != nil {
		return ch
	}
	ch := make(chan DeploymentParam)
	deploymentsChan[projectId] = ch
	go func() {
		for param := range ch {
			StartDeployment(param.Deployment, param.Commands)
		}
	}()
	return ch
}

type DeploymentParam struct {
	Deployment *models.Deployment
	Commands   contracts.Collection[*models.Command]
}

type DeploymentDetail struct {
	Deployment  *models.Deployment `json:"deployment"`
	Key         *models.Key        `json:"key"`
	RepoAddress string             `json:"repo_address"`
	ProjectPath string             `json:"project_path"`
	TimeVersion string             `json:"time_version"`
}

func CreateDeployment(project *models.Project, version, comment string, params map[string]bool, environmentsParam []int) (*models.Deployment, error) {
	results := make([]models.CommandResult, 0)
	commands := models.Commands().Where("project_id", project.Id).OrderByDesc("sort").Get()
	servers := make([]models.Server, 0)
	models.ProjectEnvironments().
		Where("project_id", project.Id).
		WhereIn("id", environmentsParam).
		Get().Foreach(func(i int, environment *models.ProjectEnvironment) {
		for _, server := range environment.Settings.Servers {
			if !server.Disabled {
				server.Environment = environment.Id
				servers = append(servers, server)
			}
		}

		models.Cabinets().WhereIn("id", environment.Settings.Cabinets).
			Get().Foreach(func(i int, cabinet *models.Cabinet) {
			for _, server := range cabinet.Settings {
				if !server.Disabled {
					server.Environment = environment.Id
					servers = append(servers, server)
				}
			}
		})
	})

	results = append(results, models.CommandResult{Step: models.Init, Servers: makeCommandOutputs(servers)})

	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.BeforeClone).ToArray(), params, servers)...)
	results = append(results, models.CommandResult{Step: models.Clone, Servers: makeCommandOutputs(servers)})
	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.AfterClone).ToArray(), params, servers)...)

	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.BeforePrepare).ToArray(), params, servers)...)
	results = append(results, models.CommandResult{Step: models.Prepare, Servers: makeCommandOutputs(servers)})
	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.AfterPrepare).ToArray(), params, servers)...)

	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.BeforeRelease).ToArray(), params, servers)...)
	results = append(results, models.CommandResult{Step: models.Release, Servers: makeCommandOutputs(servers)})
	results = append(results, makeCommandOutputsWithParams(commands.Where("step", models.AfterRelease).ToArray(), params, servers)...)

	deployment := models.Deployments().Create(contracts.Fields{
		"project_id":   project.Id,
		"params":       params,
		"results":      results,
		"environments": environmentsParam,
		"version":      version,
		"comment":      comment,
		"status":       models.StatusWaiting,
	})

	go GoDeployment(deployment, commands)

	return deployment, nil
}

func GoDeployment(deployment *models.Deployment, commands contracts.Collection[*models.Command]) {
	deploymentChan(deployment.ProjectId) <- DeploymentParam{Deployment: deployment, Commands: commands}
}

func StartDeployment(deployment *models.Deployment, commands contracts.Collection[*models.Command]) {
	project := models.Projects().FindOrFail(deployment.ProjectId)

	detail := DeploymentDetail{
		Deployment:  deployment,
		Key:         models.Keys().FindOrFail(project.KeyId),
		RepoAddress: project.RepoAddress,
		ProjectPath: project.ProjectPath,
		TimeVersion: carbon.Parse(deployment.CreatedAt).ToShortDateTimeString(),
	}

	commandsList := map[string]deploymentCommand{
		models.Init:    scriptFunc(fmt.Sprintf("mkdir -p %s/releases/%s", detail.ProjectPath, detail.TimeVersion), false),
		models.Clone:   clone,
		models.Prepare: prepare,
		models.Release: release,
	}
	commands.Foreach(func(i int, command *models.Command) {
		commandsList[fmt.Sprintf("%d", command.Id)] = scriptFunc(command.Script, true)
	})

	_ = deployment.Update(contracts.Fields{"status": models.StatusRunning})
	DeploymentNotify(deployment)
	for i, result := range deployment.Results {
		command := commandsList[result.Step]
		if result.Command > 0 {
			command = commandsList[fmt.Sprintf("%d", result.Command)]
		}

		for s, server := range result.Servers {
			server.Status = models.StatusRunning
			result.Servers[s] = server
			now := time.Now()
			deployment.Results[i] = result
			_ = deployment.Update(contracts.Fields{"results": deployment.Results})
			DeploymentNotify(deployment)

			output, err := command(detail, server.Server, "")
			server.Status = models.StatusFinished
			if err != nil {
				server.Status = models.StatusFailed
			}

			server.Outputs = output
			result.Servers[s] = server
			result.TimeConsuming = int(time.Now().Sub(now).Milliseconds())
			deployment.Results[i] = result
			_ = deployment.Update(contracts.Fields{"results": deployment.Results})
			DeploymentNotify(deployment)

			if err != nil {
				_ = deployment.Update(contracts.Fields{"status": models.StatusFailed})
				DeploymentNotify(deployment)
				return
			}
		}
	}

	_ = deployment.Update(contracts.Fields{"status": models.StatusFinished})
	DeploymentNotify(deployment)
}

func log(content string) string {
	return carbon.Now().ToDateTimeString() + "\t" + content
}

func makeCommandOutputsWithParams(commands []*models.Command, params map[string]bool, servers []models.Server) []models.CommandResult {
	results := make([]models.CommandResult, 0)
	for _, command := range commands {
		selected, existsParam := params[fmt.Sprintf("%d", command.Id)]
		if !command.Optional || selected || (!existsParam && command.DefaultSelected) {
			results = append(results, models.CommandResult{
				Step:    command.Step,
				Command: command.Id,
				Servers: makeCommandOutputs(
					collection.New(servers).Filter(func(i int, server models.Server) bool {
						return utils2.IsInT(server.Environment, command.Environments)
					}).ToArray(),
				),
			})
		}
	}
	return results
}
func makeCommandOutputs(servers []models.Server) map[string]models.CommandOutput {
	commandServers := make(map[string]models.CommandOutput)
	for _, server := range servers {
		commandServers[server.Host] = models.CommandOutput{Server: server, Status: models.StatusWaiting}
	}
	return commandServers
}

func clone(deployment DeploymentDetail, server models.Server, script string) (string, error) {
	var outputs = []string{
		log("Clone step started"),
	}

	repoPath := fmt.Sprintf("%s%s", tempRepoPath, deployment.TimeVersion+filepath.Base(deployment.ProjectPath))

	// 克隆代码到本地
	if commit, comment, err := utils.CloneRepo(
		deployment.RepoAddress,
		deployment.Key.PrivateKey,
		deployment.Deployment.Version,
		repoPath,
	); err != nil {
		outputs = append(outputs, log("Failed to clone code to piplin"))
		outputs = append(outputs, log(err.Error()))
		return strings.Join(outputs, "\n"), err
	} else {
		newAttributes := contracts.Fields{
			"commit": commit,
		}
		if deployment.Deployment.Comment == "" {
			newAttributes["comment"] = comment
		}
		err = deployment.Deployment.Update(newAttributes)
		if err != nil {
			outputs = append(outputs, log("Failed to update commit"))
		}
		outputs = append(outputs, log("The commit hash is "+commit))
		outputs = append(outputs, log("The comment is "+comment))
		outputs = append(outputs, log("Successfully clone code to piplin"))
	}

	if err := os.RemoveAll(repoPath + "/.git"); err != nil {
		outputs = append(outputs, log("Failed to remove dir .git"))
		outputs = append(outputs, log(err.Error()))
		return strings.Join(outputs, "\n"), err
	}

	zipFile := repoPath + ".zip"

	zipErr := utils.ZipFolder(repoPath, zipFile)
	if zipErr != nil {
		outputs = append(outputs, log("Failed to zip the repo folder"))
		outputs = append(outputs, log(zipErr.Error()))
		return strings.Join(outputs, "\n"), zipErr
	} else {
		_ = os.RemoveAll(repoPath)
		outputs = append(outputs, log("Successfully zip the folder"))
	}

	sftpClient, err := utils.ConnectSFTP(
		fmt.Sprintf("%s:%d", server.Host, server.Port), server.User, deployment.Key.PrivateKey,
	)
	if err != nil {
		outputs = append(outputs, log("Failed to connect server via sftp"))
		outputs = append(outputs, log(err.Error()))
		return strings.Join(outputs, "\n"), err
	} else {
		outputs = append(outputs, log("Successfully connect to server via sftp"))
	}

	// 同步代码到服务器
	err = utils.SyncFile(sftpClient, zipFile, fmt.Sprintf("%s/releases/%s.zip", deployment.ProjectPath, deployment.TimeVersion))
	if err != nil {
		outputs = append(outputs, log("Failed to sync zip to server via sftp"))
		outputs = append(outputs, log(err.Error()))
		return strings.Join(outputs, "\n"), err
	} else {
		_ = os.RemoveAll(zipFile)
		outputs = append(outputs, log("Successfully synchronisation of zip file to the server"))
	}

	var scripts = []string{
		script,
		fmt.Sprintf("cd %s/releases", deployment.ProjectPath),
		fmt.Sprintf("unzip %s.zip -d %s", deployment.TimeVersion, deployment.TimeVersion),
		fmt.Sprintf("rm %s.zip", deployment.TimeVersion),
		"echo 'successfully remove the zip file'",
	}

	output, execErr := _connectAndExec(deployment, server, scripts...)
	outputs = append(outputs, output)
	err = execErr

	return strings.Join(outputs, "\n"), err
}

func prepare(deployment DeploymentDetail, server models.Server, script string) (string, error) {
	var outputs = []string{
		log("Prepare step started"),
	}
	var inputs = []string{fmt.Sprintf("cd %s/releases/%s", deployment.ProjectPath, deployment.TimeVersion)}

	// 准备所有配置文件 start
	configFiles := models.ConfigFiles().Where("project_id", deployment.Deployment.ProjectId).Get().ToArray()
	for _, file := range configFiles {
		if utils2.IsInT(server.Environment, file.Environments) {
			inputs = append(inputs,
				fmt.Sprintf("echo '%s' > %s", file.Content, file.Path),
				"echo \"$(date '+%Y-%m-%d %H:%M:%S')\\t config file ["+file.Name+"] are prepared in "+file.Path+"\"",
			)
		}
	}
	// 准备所有配置文件 end

	// 准备所有共享目录 start
	shares := models.ShareFiles().Where("project_id", deployment.Deployment.ProjectId).Get().ToArray()
	for _, share := range shares {
		if strings.HasSuffix(share.Path, "/") {
			path := strings.TrimSuffix(share.Path, "/")
			inputs = append(inputs,
				fmt.Sprintf("mkdir -p %s/shared/%s", deployment.ProjectPath, path),
				fmt.Sprintf("cp -ruv %s/releases/%s/%s/* %s/shared/%s", deployment.ProjectPath, deployment.TimeVersion, path, deployment.ProjectPath, path),
				fmt.Sprintf("rm -rf %s/releases/%s/%s", deployment.ProjectPath, deployment.TimeVersion, path),
				fmt.Sprintf("ln -s %s/shared/%s %s/releases/%s/%s", deployment.ProjectPath, path, deployment.ProjectPath, deployment.TimeVersion, path),
				"echo \"$(date '+%Y-%m-%d %H:%M:%S')\\t shared directory ["+share.Name+"] are prepared in "+path+"\"",
			)
		} else {
			inputs = append(inputs,
				fmt.Sprintf("cp -ruv %s/releases/%s/%s %s/shared/%s", deployment.ProjectPath, deployment.TimeVersion, share.Path, deployment.ProjectPath, share.Path),
				fmt.Sprintf("rm -rf %s/releases/%s/%s", deployment.ProjectPath, deployment.TimeVersion, share.Path),
				fmt.Sprintf("ln -s %s/shared/%s %s/releases/%s/%s", deployment.ProjectPath, share.Path, deployment.ProjectPath, deployment.TimeVersion, share.Path),
				"echo \"$(date '+%Y-%m-%d %H:%M:%S')\\t shared file ["+share.Name+"] are prepared in "+share.Path+"\"",
			)
		}

	}
	// 准备所有共享目录 end

	inputs = append(inputs, script)

	output, err := _connectAndExec(deployment, server, inputs...)
	outputs = append(outputs, output)
	outputs = append(outputs, log("Preparatory steps completed"))
	return strings.Join(outputs, "\n"), err
}

func _connectAndExec(deployment DeploymentDetail, server models.Server, script ...string) (output string, err error) {
	var outputs []string
	if len(script) == 0 {
		return
	}
	defer func() {
		output = strings.Join(outputs, "\n")
	}()

	client, err := utils.ConnectToSSHServer(
		fmt.Sprintf("%s:%d", server.Host, server.Port),
		deployment.Key.PrivateKey,
		server.User,
	)
	if err != nil {
		outputs = append(outputs, log("Failed to connect the server"))
		outputs = append(outputs, log(err.Error()))
		return strings.Join(outputs, "\n"), err
	} else {
		outputs = append(outputs, log("Successfully connect to the server"))
	}

	outputs = append(outputs, log("Script started"))
	execOutput, err := utils.ExecuteSSHCommand(client, script...)
	if execOutput != "" {
		outputs = append(outputs, execOutput)
	}
	if err != nil {
		outputs = append(outputs, log("Failed to exec the script"))
		outputs = append(outputs, log(err.Error()))
	} else {
		outputs = append(outputs, log("Script execution complete"))
	}
	return
}

func release(deployment DeploymentDetail, server models.Server, script string) (string, error) {
	var outputs []string
	var inputs = []string{
		fmt.Sprintf("rm %s/current", deployment.ProjectPath),
		"echo 'remove old link'",
		fmt.Sprintf("ln -s %s/releases/%s %s/current", deployment.ProjectPath, deployment.TimeVersion, deployment.ProjectPath),
		"echo 'successfully create a new link'",
	}

	inputs = append(inputs, script)

	output, err := _connectAndExec(deployment, server, inputs...)
	outputs = append(outputs, output)
	return strings.Join(outputs, "\n"), err
}

func scriptFunc(script string, cd bool) deploymentCommand {
	return func(deployment DeploymentDetail, server models.Server, _ string) (string, error) {
		var outputs = []string{
			log("script started"),
		}

		var inputsScript = []string{script}
		if cd {
			inputsScript = []string{fmt.Sprintf("cd %s/releases/%s", deployment.ProjectPath, deployment.TimeVersion), script}
		}

		output, err := _connectAndExec(deployment, server, inputsScript...)
		outputs = append(outputs, output)

		return strings.Join(outputs, "\n"), err
	}
}
