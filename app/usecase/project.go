package usecase

import (
	"errors"
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
		key, err = CreateKey(utils.ToString(fields["name"], ""))
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
