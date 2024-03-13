package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateConfig(projectId int, fields contracts.Fields) (*models.ConfigFile, error) {
	if models.ConfigFiles().
		Where("project_id", projectId).
		Where("name", fields["name"]).
		Count() > 0 {
		return nil, errors.New("配置文件已存在！")
	}
	fields["project_id"] = projectId
	return models.ConfigFiles().CreateE(
		utils.OnlyFields(fields, "name", "project_id", "path", "content", "environments"),
	)
}

func UpdateConfig(id any, fields contracts.Fields) error {
	env := models.ConfigFiles().Find(id)
	if models.ConfigFiles().
		Where("project_id", env.ProjectId).
		Where("id", "!=", id).
		Where("name", fields["name"]).
		Count() > 0 {
		return errors.New("配置文件已存在！")
	}

	_, err := models.ConfigFiles().Where("id", id).
		UpdateE(utils.OnlyFields(fields, "content", "name", "path", "environments"))

	return err
}

func DeleteConfig(id any) error {
	_, err := models.ConfigFiles().WhereIn("id", id).DeleteE()
	return err
}
