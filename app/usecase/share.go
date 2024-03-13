package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateShare(projectId int, fields contracts.Fields) (*models.ShareFile, error) {
	if models.ShareFiles().
		Where("project_id", projectId).
		Where("name", fields["name"]).
		Count() > 0 {
		return nil, errors.New("共享文件已存在！")
	}
	fields["project_id"] = projectId
	return models.ShareFiles().CreateE(
		utils.OnlyFields(fields, "name", "project_id", "path"),
	)
}

func UpdateShare(id any, fields contracts.Fields) error {
	env := models.ShareFiles().Find(id)
	if models.ShareFiles().
		Where("project_id", env.ProjectId).
		Where("id", "!=", id).
		Where("name", fields["name"]).
		Count() > 0 {
		return errors.New("共享文件已存在！")
	}

	_, err := models.ShareFiles().Where("id", id).
		UpdateE(utils.OnlyFields(fields, "name", "path"))

	return err
}

func DeleteShare(id any) error {
	_, err := models.ShareFiles().WhereIn("id", id).DeleteE()
	return err
}
