package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
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
