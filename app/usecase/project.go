package usecase

import (
	"errors"
	"github.com/goal-web/collection"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
	utils2 "github.com/qbhy/goal-piplin/app/utils"
	"github.com/savsgio/gotils/uuid"
)

func CreateProject(fields contracts.Fields) (models.Project, error) {
	fields = utils.OnlyFields(fields,
		"name", "default_branch", "project_path", "repo_address", "group_id", "key_id", "creator_id")
	var project models.Project
	var key models.Key
	var err error

	if models.Projects().Where("name", fields["name"]).Count() > 0 {
		return project, errors.New("项目已存在")
	}

	var existsKey = utils.ToInt(fields["key_id"], 0) > 0
	if !existsKey {
		key, err = CreateKey(project.CreatorId, utils.ToString(fields["name"], ""))
		fields["key_id"] = key.Id
		if err != nil {
			return project, err
		}
	}

	fields["uuid"] = uuid.V4()
	fields["settings"] = models.ProjectSettings{}
	project = models.Projects().Create(fields)

	if existsKey {
		err = UpdateProjectBranches(project, key)
	}

	return project, err
}

func CopyProject(targetProject models.Project, fields contracts.Fields) (models.Project, error) {
	targetKey := models.Keys().FindOrFail(targetProject.KeyId)
	key := models.Keys().Create(contracts.Fields{
		"creator_id":  fields["creator_id"],
		"name":        fields["name"],
		"public_key":  targetKey.PublicKey,
		"private_key": targetKey.PrivateKey,
	})
	utils.MergeFields(fields, contracts.Fields{
		"uuid":         uuid.V4(),
		"project_path": targetProject.ProjectPath,
		"key_id":       key.Id,
		"settings":     targetProject.Settings,
	})

	project := models.Projects().Create(fields)
	environmentsMap := make(map[int]int)

	models.ProjectEnvironments().Where("project_id", targetProject.Id).Get().Foreach(func(i int, env models.ProjectEnvironment) {
		environmentsMap[env.Id] = models.ProjectEnvironments().Create(contracts.Fields{
			"project_id": project.Id,
			"name":       env.Name,
			"settings":   env.Settings,
		}).Id
	})

	models.Commands().Where("project_id", targetProject.Id).Get().Foreach(func(i int, command models.Command) {
		models.Commands().Create(contracts.Fields{
			"name":       command.Name,
			"project_id": project.Id,
			"step":       command.Step,
			"sort":       command.Sort,
			"user":       command.User,
			"script":     command.Script,
			"environments": collection.New(command.Environments).Each(func(_ int, envId int) int {
				return environmentsMap[envId]
			}).ToArray(),
			"optional":         command.Optional,
			"default_selected": command.DefaultSelected,
		})
	})

	models.ConfigFiles().Where("project_id", targetProject.Id).Get().Foreach(func(i int, config models.ConfigFile) {
		models.ConfigFiles().Create(contracts.Fields{
			"project_id": project.Id,
			"name":       config.Name,
			"path":       config.Path,
			"content":    config.Content,
			"environments": collection.New(config.Environments).Each(func(_ int, envId int) int {
				return environmentsMap[envId]
			}).ToArray(),
		})
	})

	models.ShareFiles().Where("project_id", targetProject.Id).Get().Foreach(func(i int, share models.ShareFile) {
		models.ShareFiles().Create(contracts.Fields{
			"project_id": project.Id,
			"name":       share.Name,
			"path":       share.Path,
		})
	})

	return project, nil
}

func UpdateProjectBranches(project models.Project, key models.Key) error {
	branches, tags, err := GetBranchDetail(project, key)
	if err == nil {
		project.Settings = models.ProjectSettings{
			Branches:  branches,
			Tags:      tags,
			EnvVars:   project.Settings.EnvVars,
			Callbacks: project.Settings.Callbacks,
		}
		models.Projects().Where("id", project.Id).Update(contracts.Fields{
			"settings": project.Settings,
		})
	}
	return err
}

func UpdateProject(id int, fields contracts.Fields) (models.Project, error) {
	project := models.Projects().FindOrFail(id)
	if models.Projects().Where("id", "!=", id).Where("name", fields["name"]).Count() > 0 {
		return project, errors.New("项目已存在")
	}
	_, err := models.Projects().Where("id", id).UpdateE(utils.OnlyFields(fields,
		"name", "default_branch", "project_path", "repo_address", "group_id", "key_id",
	))

	if err == nil {
		return project, UpdateProjectBranches(project, models.Keys().FindOrFail(project.KeyId))
	}

	return models.Projects().FindOrFail(id), err
}

func GetProjectDetail(id any) models.ProjectDetail {
	project := models.Projects().Find(id)
	return models.ProjectDetail{
		Project: project,
		Key:     models.Keys().Find(project.KeyId),
		Group:   models.Groups().Find(project.GroupId),
		Members: table.ArrayQuery("user_projects").
			Select("user_id", "username", "nickname", "avatar", "status", "user_projects.id").
			Where("project_id", project.Id).
			LeftJoin("users", "users.id", "=", "user_projects.user_id").
			Get().ToArrayFields(),
	}
}

func GetBranchDetail(project models.Project, key models.Key) ([]string, []string, error) {
	return utils2.GetRepositoryBranchesAndTags(project.RepoAddress, key.PrivateKey)
}

func DeleteProject(project models.Project) error {

	if models.Projects().Where("key_id", project.KeyId).Count() == 1 {
		models.Keys().Where("id", project.KeyId).Delete()
	}

	_, err := models.Projects().WhereIn("id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.ConfigFiles().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.UserProjects().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.ShareFiles().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.ProjectEnvironments().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.ProjectEnvironments().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.Deployments().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	_, err = models.Commands().WhereIn("project_id", project.Id).DeleteE()
	if err != nil {
		return err
	}

	return err
}

// HasProjectPermission 判断用户是否存在指定项目的权限
func HasProjectPermission(project models.Project, userId int) bool {
	if project.CreatorId == userId {
		return true
	}

	if project.GroupId > 0 && HasGroupPermission(models.Groups().FindOrFail(project.GroupId), userId) {
		return true
	}

	return models.UserProjects().
		Where("project_id", project.Id).
		Where("user_id", userId).
		Where("status", models.InviteStatusJoined).
		Count() > 0
}
