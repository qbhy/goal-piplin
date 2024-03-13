package usecase

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateCommand(projectId int, fields contracts.Fields) (*models.Command, error) {
	fields["project_id"] = projectId
	return models.Commands().CreateE(
		utils.OnlyFields(fields, "name", "project_id", "step", "sort", "user", "script", "environments", "optional", "default_selected"),
	)
}

func UpdateCommand(id any, fields contracts.Fields) error {
	_, err := models.Commands().Where("id", id).
		UpdateE(utils.OnlyFields(fields, "name", "step", "sort", "user", "script", "environments", "optional", "default_selected"))

	return err
}

func DeleteCommand(id any) error {
	_, err := models.Commands().WhereIn("id", id).DeleteE()
	return err
}
