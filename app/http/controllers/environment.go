package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetEnvironments(request contracts.HttpRequest) any {
	list, total := models.ProjectEnvironments().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.ProjectEnvironment]) contracts.Query[models.ProjectEnvironment] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
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

func UpdateEnvironment(request contracts.HttpRequest) any {
	err := usecase.UpdateEnvironment(request.Get("id"), request.GetString("name"), request.Get("settings"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteEnvironment(request contracts.HttpRequest) any {
	err := usecase.DeleteEnvironment(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
