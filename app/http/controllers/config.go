package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetConfigs(request contracts.HttpRequest) any {
	list, total := models.ConfigFiles().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.ConfigFile]) contracts.Query[models.ConfigFile] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(request.GetString("content") != "", func(q contracts.QueryBuilder[models.ConfigFile]) contracts.Query[models.ConfigFile] {
			return q.Where("content", "like", "%"+request.GetString("content")+"%")
		}).
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateConfig(request contracts.HttpRequest) any {

	Config, err := usecase.CreateConfig(request.GetInt("project_id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{
		"data": Config,
	}
}

func UpdateConfig(request contracts.HttpRequest) any {
	err := usecase.UpdateConfig(request.Get("id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteConfig(request contracts.HttpRequest) any {
	err := usecase.DeleteConfig(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
