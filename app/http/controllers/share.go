package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetShares(request contracts.HttpRequest, guard contracts.Guard) any {
	targetProject := models.Projects().FindOrFail(request.QueryParam("project_id"))
	if !usecase.HasProjectPermission(targetProject, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	list, total := models.ShareFiles().
		Where("project_id", targetProject.Id).
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.Query[models.ShareFile]) contracts.Query[models.ShareFile] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateShare(request contracts.HttpRequest, guard contracts.Guard) any {
	targetProject := models.Projects().FindOrFail(request.GetInt("project_id"))
	if !usecase.HasProjectPermission(targetProject, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	Share, err := usecase.CreateShare(targetProject.Id, request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{
		"data": Share,
	}
}

func UpdateShare(request contracts.HttpRequest, guard contracts.Guard) any {
	share := models.ShareFiles().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(share.ProjectId)
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	err := usecase.UpdateShare(share.Id, request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteShare(request contracts.HttpRequest, guard contracts.Guard) any {
	share := models.ShareFiles().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(share.ProjectId)
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	err := usecase.DeleteShare(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
