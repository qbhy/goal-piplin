package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetConfigs(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.QueryParam("project_id"))

	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	list, total := models.ConfigFiles().
		Where("project_id", project.Id).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.ConfigFile]) contracts.Query[models.ConfigFile] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(request.GetString("content") != "", func(q contracts.QueryBuilder[models.ConfigFile]) contracts.Query[models.ConfigFile] {
			return q.Where("content", "like", "%"+request.GetString("content")+"%")
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
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

func UpdateConfig(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.Get("id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	err := usecase.UpdateConfig(project.Id, request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteConfig(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.Get("id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	err := usecase.DeleteConfig(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
