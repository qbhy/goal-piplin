package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetEnvironments(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.QueryParam("project_id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	list, total := models.ProjectEnvironments().
		Where("project_id", project.Id).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.ProjectEnvironment]) contracts.Query[models.ProjectEnvironment] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateEnvironment(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.Get("project_id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	environment, err := usecase.CreateEnvironment(request.GetString("name"), project.Id)

	if err != nil {
		return contracts.Fields{"msg": "创建环境失败"}
	}

	return contracts.Fields{
		"data": environment,
	}
}

func UpdateEnvironment(request contracts.HttpRequest, guard contracts.Guard) any {
	env := models.ProjectEnvironments().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(env.ProjectId)
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	err := usecase.UpdateEnvironment(env.Id, request.GetString("name"), request.Get("settings"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteEnvironment(request contracts.HttpRequest, guard contracts.Guard) any {
	env := models.ProjectEnvironments().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(env.ProjectId)
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	err := usecase.DeleteEnvironment(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
