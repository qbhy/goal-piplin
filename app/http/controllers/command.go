package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetCommands(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.QueryParam("project_id"))
	if !usecase.HasProjectPermission(project, utils.ToInt(guard.GetId(), 0)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	list, total := models.Commands().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.Command]) contracts.Query[models.Command] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(request.GetString("user") != "", func(q contracts.QueryBuilder[models.Command]) contracts.Query[models.Command] {
			return q.Where("user", "like", "%"+request.GetString("user")+"%")
		}).
		When(request.GetString("step") != "", func(q contracts.QueryBuilder[models.Command]) contracts.Query[models.Command] {
			return q.Where("step", request.GetString("step"))
		}).
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateCommand(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.Get("project_id"))
	if !usecase.HasProjectPermission(project, utils.ToInt(guard.GetId(), 0)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	Command, err := usecase.CreateCommand(project.Id, request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{
		"data": Command,
	}
}

func UpdateCommand(request contracts.HttpRequest, guard contracts.Guard) any {
	command := models.Commands().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(command.ProjectId)
	if !usecase.HasProjectPermission(project, utils.ToInt(guard.GetId(), 0)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	err := usecase.UpdateCommand(command.Id, request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteCommand(request contracts.HttpRequest, guard contracts.Guard) any {
	command := models.Commands().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(command.ProjectId)
	if !usecase.HasProjectPermission(project, utils.ToInt(guard.GetId(), 0)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	err := usecase.DeleteCommand(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
