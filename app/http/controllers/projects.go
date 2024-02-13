package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/requests"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/example/app/usecase"
	"github.com/goal-web/validation"
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
		"data": usecase.GetProjectDetail(request.Param("id")),
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
