package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateUserProject(projectId, userId int) (*models.UserProject, error) {
	if models.UserProjects().
		Where("project_id", projectId).
		Where("user_id", userId).
		Count() > 0 {
		return nil, errors.New("邀请失败，您已邀请该用户")
	}
	return models.UserProjects().CreateE(contracts.Fields{
		"user_id":    userId,
		"project_id": projectId,
		"status":     models.InviteStatusWaiting,
	})
}

func UpdateUserProject(project *models.UserProject, status string) error {
	_, err := models.UserProjects().Where("id", project.Id).UpdateE(contracts.Fields{
		"status": status,
	})
	return err
}

func DeleteUserProject(id any) error {
	_, err := models.UserProjects().WhereIn("id", id).DeleteE()
	return err
}
