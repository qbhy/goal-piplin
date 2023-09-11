package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func GetProjects(request contracts.HttpRequest) any {
	list, total := models.Projects().OrderByDesc("id").Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"list":  list,
	}
}
