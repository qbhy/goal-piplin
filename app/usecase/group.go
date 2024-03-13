package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
)

func DeleteGroups(id any) error {
	_, err := models.Groups().WhereIn("id", id).DeleteE()
	return err
}

func UpdateGroup(id any, fields contracts.Fields) error {
	if models.Groups().Where("id", "!=", id).Where("name", fields["name"]).Count() > 0 {
		return errors.New("分组名称")
	}

	_, err := models.Groups().Where("id", id).UpdateE(utils.OnlyFields(fields, "name"))

	return err
}
