package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/requests"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/validation"
	"github.com/savsgio/gotils/uuid"
	"time"
)

func GetProjects(request contracts.HttpRequest) any {
	list, total := models.Projects().OrderByDesc("id").Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"list":  list.ToArray(),
	}
}

func CreateProject(request requests.ProjectRequest, guard contracts.Guard) any {
	validation.VerifyForm(request)
	fields := request.Fields()
	fields["uuid"] = uuid.V4()
	fields["creator_id"] = guard.GetId()
	fields["created_at"] = time.Now()
	return models.Projects().Create(fields)
}
