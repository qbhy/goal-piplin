package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/validation"
	"github.com/qbhy/goal-piplin/app/http/requests"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetProjects(request contracts.HttpRequest) any {
	list, total := models.Projects().OrderByDesc("id").Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func GetProject(request contracts.HttpRequest) any {
	return contracts.Fields{
		"data": usecase.GetProjectDetail(request.GetString("id")),
	}
}

func CreateProject(request requests.ProjectRequest, guard contracts.Guard) any {
	validation.VerifyForm(request)
	fields := request.Fields()
	fields["creator_id"] = guard.GetId()

	project, err := usecase.CreateProject(fields)

	if err != nil {
		return contracts.Fields{"msg": "创建项目失败"}
	}

	return contracts.Fields{
		"data": project,
	}
}

func UpdateProject(request requests.ProjectRequest) any {
	validation.VerifyForm(request)
	fields := request.Fields()

	if err := usecase.UpdateProject(request.GetInt("id"), fields); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
