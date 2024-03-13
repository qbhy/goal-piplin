package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetCommands(request contracts.HttpRequest) any {
	list, total := models.Commands().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateCommand(request contracts.HttpRequest) any {

	Command, err := usecase.CreateCommand(request.GetInt("project_id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{
		"data": Command,
	}
}

func UpdateCommand(request contracts.HttpRequest) any {
	err := usecase.UpdateCommand(request.Get("id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteCommand(request contracts.HttpRequest) any {
	err := usecase.DeleteCommand(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
