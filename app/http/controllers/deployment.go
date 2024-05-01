package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/http/requests"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetDeployments(request contracts.HttpRequest) any {
	list, total := models.Deployments().
		Where("project_id", request.Get("project_id")).
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func GetDeploymentDetail(request contracts.HttpRequest) any {
	deployment := models.Deployments().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(deployment.ProjectId)

	return contracts.Fields{
		"data": contracts.Fields{
			"deployment": deployment,
			"project":    project,
			"commands":   models.Commands().Where("project_id", project.Id).Get().ToArray(),
		},
	}
}

func CreateDeployment(request contracts.HttpRequest) any {
	var form requests.DeploymentRequest
	if err := request.Parse(&form); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	project := models.Projects().FindOrFail(form.ProjectId)
	if deployment, err := usecase.CreateDeployment(project, form.Version, form.Comment, form.Params, form.Environments); err != nil {
		return contracts.Fields{"msg": err.Error()}
	} else {
		return contracts.Fields{"data": deployment}
	}
}

func Notify(sse contracts.SseFactory, request contracts.HttpRequest) any {
	var msg string
	if err := sse.Sse("/api/notify").Broadcast(models.Deployments().FindOrFail(request.Get("id"))); err != nil {
		msg = err.Error()
	}
	return contracts.Fields{
		"err": msg,
	}
}
