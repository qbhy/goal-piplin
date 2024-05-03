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

func init() {
	tempRepoPath = os.TempDir()
}

type DeploymentDetail struct {
	Key         models.Key `json:"key"`
	ProjectId   int        `json:"project_id"`
	Version     string     `json:"version"`
	RepoAddress string     `json:"repo_address"`
	ProjectPath string     `json:"project_path"`
	TimeVersion string     `json:"time_version"`
}

func CreateDeployment(project models.Project, version, comment string, params map[string]bool, environmentsParam []int) (models.Deployment, error) {
	results := make([]models.CommandResult, 0)
	commands := models.Commands().Where("project_id", project.Id).OrderByDesc("sort").Get()
	servers := make([]models.Server, 0)
	models.ProjectEnvironments().
		Where("project_id", project.Id).
		WhereIn("id", environmentsParam).
		Get().Foreach(func(i int, environment models.ProjectEnvironment) {
		for _, server := range environment.Settings.Servers {
			if !server.Disabled {
				server.Environment = environment.Id
				servers = append(servers, server)
			}
		}

		models.Cabinets().WhereIn("id", environment.Settings.Cabinets).
			Get().Foreach(func(i int, cabinet models.Cabinet) {
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

	go StartDeployment(deployment, commands)

	return deployment, nil
}

func StartDeployment(deployment models.Deployment, commands contracts.Collection[models.Command]) {
	project := models.Projects().FindOrFail(deployment.ProjectId)

	detail := DeploymentDetail{
		ProjectId:   deployment.ProjectId,
		Key:         models.Keys().FindOrFail(project.KeyId),
		Version:     deployment.Version,
		RepoAddress: project.RepoAddress,
		ProjectPath: project.ProjectPath,
		TimeVersion: carbon.Parse(deployment.CreatedAt).ToShortDateTimeString(),
	}

	commandsList := map[string]deploymentCommand{
		models.Init:    scriptFunc(fmt.Sprintf("mkdir -p %s/releases/%s", detail.ProjectPath, detail.TimeVersion)),
		models.Clone:   clone,
		models.Prepare: prepare,
		models.Release: release,
	}
	commands.Foreach(func(i int, command models.Command) {
		commandsList[fmt.Sprintf("%d", command.Id)] = scriptFunc(command.Script)
	})

	models.Deployments().Where("id", deployment.Id).Update(contracts.Fields{
		"status": models.StatusRunning,
	})
	deployment.Status = models.StatusRunning
	DeploymentNotify(deployment)
	// todo 锁
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
			models.Deployments().Where("id", deployment.Id).Update(contracts.Fields{
				"results": deployment.Results,
			})
			DeploymentNotify(deployment)

			output, err := command(detail, server.Server, "")
			server.Status = models.StatusFinished
			if err != nil {
				server.Status = models.StatusFailed
				output += err.Error()
			}

			server.Outputs = output
			result.Servers[s] = server
			result.TimeConsuming = int(time.Now().Sub(now).Milliseconds())
			deployment.Results[i] = result
			models.Deployments().Where("id", deployment.Id).Update(contracts.Fields{
				"results": deployment.Results,
			})
			DeploymentNotify(deployment)

			if err != nil {
				models.Deployments().Where("id", deployment.Id).Update(contracts.Fields{
					"status": models.StatusFailed,
				})
				deployment.Status = models.StatusFailed
				DeploymentNotify(deployment)
				return
			}
		}
	}

	models.Deployments().Where("id", deployment.Id).Update(contracts.Fields{
		"status": models.StatusFinished,
	})
	deployment.Status = models.StatusFinished
	DeploymentNotify(deployment)
}

func makeCommandOutputsWithParams(commands []models.Command, params map[string]bool, servers []models.Server) []models.CommandResult {
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
	var outputs []string

	client, err := utils.ConnectToSSHServer(
		fmt.Sprintf("%s:%d", server.Host, server.Port),
		deployment.Key.PrivateKey,
		server.User,
	)
	if err != nil {
		return "", err
	}

	repoPath := fmt.Sprintf("%s%s", tempRepoPath, deployment.TimeVersion+filepath.Base(deployment.ProjectPath))

	// 克隆代码到本地
	if err = utils.CloneRepoBranchOrCommit(
		deployment.RepoAddress,
		deployment.Key.PrivateKey,
		deployment.Version,
		repoPath,
	); err != nil {
		return "", err
	}

	if err = os.RemoveAll(repoPath + "/.git"); err != nil {
		return "", err
	}

	sftpClient, err := utils.ConnectSFTP(
		fmt.Sprintf("%s:%d", server.Host, server.Port), server.User, deployment.Key.PrivateKey,
	)
	if err != nil {
		return strings.Join(outputs, "\n"), err
	}

	// 同步代码到服务器
	err = utils.SyncDir(sftpClient, repoPath, fmt.Sprintf("%s/releases/%s", deployment.ProjectPath, deployment.TimeVersion))
	if err != nil {
		return strings.Join(outputs, "\n"), err
	}
	outputs = append(outputs, "File synchronization successful.")

	// 执行脚本，如果有的话
	output, err := utils.ExecuteSSHCommand(client, script)
	if err == nil && output != "" {
		outputs = append(outputs, output)
	}

	_ = os.RemoveAll(repoPath)

	return strings.Join(outputs, "\n"), err
}

func prepare(deployment DeploymentDetail, server models.Server, script string) (string, error) {
	var outputs []string
	var inputs = []string{fmt.Sprintf("cd %s/releases/%s", deployment.ProjectPath, deployment.TimeVersion)}

	// 准备所有配置文件 start
	configFiles := models.ConfigFiles().Where("project_id", deployment.ProjectId).Get().ToArray()
	for _, file := range configFiles {
		if utils2.IsInT(server.Environment, file.Environments) {
			inputs = append(inputs, fmt.Sprintf("echo '%s' > %s", file.Content, file.Path))
		}
	}
	// 准备所有配置文件 end

	// 准备所有共享目录 start
	shares := models.ShareFiles().Where("project_id", deployment.ProjectId).Get().ToArray()
	for _, share := range shares {
		inputs = append(inputs,
			fmt.Sprintf("mkdir -p %s/shared/%s", deployment.ProjectPath, share.Path),
			fmt.Sprintf("cp -ruv %s/releases/%s/%s/* %s/shared/%s", deployment.ProjectPath, deployment.TimeVersion, share.Path, deployment.ProjectPath, share.Path),
			fmt.Sprintf("rm -rf %s/releases/%s/%s", deployment.ProjectPath, deployment.TimeVersion, share.Path),
			fmt.Sprintf("ln -s %s/shared/%s %s/releases/%s/%s", deployment.ProjectPath, share.Path, deployment.ProjectPath, deployment.TimeVersion, share.Path),
		)
	}
	// 准备所有共享目录 end

	inputs = append(inputs, script)

	client, err := utils.ConnectToSSHServer(
		fmt.Sprintf("%s:%d", server.Host, server.Port),
		deployment.Key.PrivateKey,
		server.User,
	)
	if err != nil {
		return "", err
	}

	output, err := utils.ExecuteSSHCommand(client, inputs...)
	if err == nil && output != "" {
		outputs = append(outputs, output)
	}
	return strings.Join(outputs, "\n"), err
}

func release(deployment DeploymentDetail, server models.Server, script string) (string, error) {
	var outputs []string
	var inputs = []string{
		fmt.Sprintf("rm %s/current", deployment.ProjectPath),
		fmt.Sprintf("ln -s %s/releases/%s %s/current", deployment.ProjectPath, deployment.TimeVersion, deployment.ProjectPath),
	}

	inputs = append(inputs, script)

	client, err := utils.ConnectToSSHServer(
		fmt.Sprintf("%s:%d", server.Host, server.Port),
		deployment.Key.PrivateKey,
		server.User,
	)
	if err != nil {
		return "", err
	}

	output, err := utils.ExecuteSSHCommand(client, inputs...)
	if err == nil && output != "" {
		outputs = append(outputs, output)
	}
	return strings.Join(outputs, "\n"), err
}

func scriptFunc(script string) deploymentCommand {
	return func(deployment DeploymentDetail, server models.Server, _ string) (string, error) {
		var outputs []string

		client, err := utils.ConnectToSSHServer(
			fmt.Sprintf("%s:%d", server.Host, server.Port),
			deployment.Key.PrivateKey,
			server.User,
		)
		if err != nil {
			return "", err
		}

		// 执行脚本
		output, err := utils.ExecuteSSHCommand(client,
			fmt.Sprintf("cd %s/releases/%s", deployment.ProjectPath, deployment.TimeVersion),
			script)
		if err == nil && output != "" {
			outputs = append(outputs, output)
		}
		return strings.Join(outputs, "\n"), err
	}
}
