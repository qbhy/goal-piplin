package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func CreateUserProject(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.GetInt("project_id"))
	data, err := usecase.CreateUserProject(project.Id, request.GetInt("user_id"))
	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}
	return contracts.Fields{"data": data}
}

func DeleteUserProjects(request contracts.HttpRequest) any {
	err := usecase.DeleteUserProject(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func UpdateUserProject(request contracts.HttpRequest, guard contracts.Guard) any {
	userProject := models.UserProjects().
		Where("project_id", request.Get("project_id")).
		Where("user_id", guard.GetId()).FirstOrFail()

	err := usecase.UpdateUserProject(userProject, request.GetString("status"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
