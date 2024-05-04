package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateUserGroup(groupId, userId int) (*models.UserGroup, error) {
	if models.UserGroups().
		Where("group_id", groupId).
		Where("user_id", userId).
		Count() > 0 {
		return nil, errors.New("邀请失败，您已邀请该用户")
	}
	return models.UserGroups().CreateE(contracts.Fields{
		"user_id":  userId,
		"group_id": groupId,
		"status":   models.InviteStatusWaiting,
	})
}

func UpdateUserGroup(project models.UserGroup, status string) error {
	_, err := models.UserGroups().Where("id", project.Id).UpdateE(contracts.Fields{
		"status": status,
	})
	return err
}

func DeleteUserGroup(id any) error {
	_, err := models.UserGroups().WhereIn("id", id).DeleteE()
	return err
}
