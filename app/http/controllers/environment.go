package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/example/app/usecase"
)

func GetEnvironments(request contracts.HttpRequest) any {
	list, total := models.ProjectEnvironments().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateEnvironment(request contracts.HttpRequest) any {

	environment, err := usecase.CreateEnvironment(request.GetString("name"), request.GetInt("project_id"))

	if err != nil {
		return contracts.Fields{"msg": "创建环境失败"}
	}

	return contracts.Fields{
		"data": environment,
	}
}
