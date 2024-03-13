package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetShares(request contracts.HttpRequest) any {
	list, total := models.ShareFiles().
		Where("project_id", request.QueryParam("project_id")).
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateShare(request contracts.HttpRequest) any {

	Share, err := usecase.CreateShare(request.GetInt("project_id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{
		"data": Share,
	}
}

func UpdateShare(request contracts.HttpRequest) any {
	err := usecase.UpdateShare(request.Get("id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteShare(request contracts.HttpRequest) any {
	err := usecase.DeleteShare(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
