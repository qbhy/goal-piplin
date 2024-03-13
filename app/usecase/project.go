package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/savsgio/gotils/uuid"
)

func CreateProject(fields contracts.Fields) (models.Project, error) {
	fields = utils.OnlyFields(fields,
		"name", "default_branch", "project_path", "repo_address", "group_id", "key_id", "creator_id")
	var project models.Project

	if models.Projects().Where("name", fields["name"]).Count() > 0 {
		return project, errors.New("项目已存在")
	}

	if utils.ToInt(fields["key_id"], 0) == 0 {
		key, err := CreateKey(utils.ToString(fields["name"], ""))
		fields["key_id"] = key.Id
		if err != nil {
			return project, err
		}
	}

	fields["uuid"] = uuid.V4()
	fields["settings"] = models.ProjectSettings{}

	project = models.Projects().Create(fields)
	return project, nil
}

func UpdateProject(id int, fields contracts.Fields) error {
	if models.Projects().Where("id", "!=", id).Where("name", fields["name"]).Count() > 0 {
		return errors.New("项目已存在")
	}

	_, err := models.Projects().Where("id", id).UpdateE(utils.OnlyFields(fields,
		"name", "default_branch", "project_path", "repo_address", "group_id", "key_id",
	))

	return err
}

func GetProjectDetail(id any) models.ProjectDetail {
	project := models.Projects().Find(id)
	return models.ProjectDetail{
		Project: project,
		Key:     models.Keys().Find(project.KeyId),
		Group:   models.Groups().Find(project.GroupId),
	}
}
