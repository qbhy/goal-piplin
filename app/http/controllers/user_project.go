package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetUserProjects(request contracts.HttpRequest, guard contracts.Guard) any {
	user := guard.User().(models.User)
	list, total := table.ArrayQuery("user_projects").
		OrderByDesc("user_projects.id").
		Select("project_id", "projects.name as project_name", "user_projects.created_at", "user_projects.id", "status").
		LeftJoin("projects", "projects.id", "=", "user_projects.project_id").
		When(request.GetString("project_name") != "", func(q contracts.QueryBuilder[contracts.Fields]) contracts.Query[contracts.Fields] {
			return q.Where("projects.name", "like", "%"+request.GetString("project_name")+"%")
		}).
		When(request.GetString("status") != "", func(q contracts.QueryBuilder[contracts.Fields]) contracts.Query[contracts.Fields] {
			return q.Where("status", request.GetString("status"))
		}).
		Where("user_id", user.Id).
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateUserProject(request contracts.HttpRequest, guard contracts.Guard) any {
	userId := request.GetInt("user_id")
	project := models.Projects().FindOrFail(request.GetInt("project_id"))
	data, err := usecase.CreateUserProject(project.Id, userId)
	if request.GetString("user_id") == guard.GetId() {
		return contracts.Fields{"msg": "不可以邀请自己"}
	}
	if project.CreatorId == userId {
		return contracts.Fields{"msg": "不可以邀请创建者"}
	}

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