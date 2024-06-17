package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/http/requests"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func PostDeployment(request contracts.HttpRequest) any {
	var form requests.PostDeploymentRequest
	if err := request.Parse(&form); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}
	if len(form.Environments) == 0 {
		return contracts.Fields{"msg": "请选择环境"}
	}
	if form.Version == "" {
		return contracts.Fields{"msg": "请选择部署版本"}
	}

	project := models.Projects().Where("uuid", form.UUID).FirstOrFail()
	//_ = usecase.UpdateProjectBranches(project, models.Keys().FindOrFail(project.KeyId))

	if form.Params == nil {
		form.Params = make(map[string]bool)
	}

	if deployment, err := usecase.CreateDeployment(project, form.Version, form.Comment, form.Params, form.Environments); err != nil {
		return contracts.Fields{"msg": err.Error()}
	} else {
		return contracts.Fields{"data": deployment}
	}
}

func GetDeployments(request contracts.HttpRequest, guard contracts.Guard) any {

	project := models.Projects().FindOrFail(request.Get("project_id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	list, total := models.Deployments().
		Where("project_id", request.Get("project_id")).
		When(request.GetString("comment") != "", func(q contracts.Query[models.Deployment]) contracts.Query[models.Deployment] {
			return q.Where("comment", "like", "%"+request.GetString("comment")+"%")
		}).
		When(request.GetString("version") != "", func(q contracts.Query[models.Deployment]) contracts.Query[models.Deployment] {
			return q.Where("version", "like", "%"+request.GetString("version")+"%")
		}).
		When(request.GetString("status") != "", func(q contracts.Query[models.Deployment]) contracts.Query[models.Deployment] {
			return q.Where("status", request.GetString("status"))
		}).
		OrderByDesc("id").
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func GetDeploymentDetail(request contracts.HttpRequest, guard contracts.Guard) any {
	deployment := models.Deployments().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(deployment.ProjectId)

	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	return contracts.Fields{
		"data": contracts.Fields{
			"deployment": deployment,
			"project":    project,
			"commands":   models.Commands().Where("project_id", project.Id).Get().ToArray(),
		},
	}
}

func CreateDeployment(request contracts.HttpRequest, guard contracts.Guard) any {
	var form requests.DeploymentRequest
	if err := request.Parse(&form); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}
	if len(form.Environments) == 0 {
		return contracts.Fields{"msg": "请选择环境"}
	}
	if form.Version == "" {
		return contracts.Fields{"msg": "请选择部署版本"}
	}

	project := models.Projects().FindOrFail(form.ProjectId)

	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	if form.Params == nil {
		form.Params = make(map[string]bool)
	}

	if deployment, err := usecase.CreateDeployment(project, form.Version, form.Comment, form.Params, form.Environments); err != nil {
		return contracts.Fields{"msg": err.Error()}
	} else {
		return contracts.Fields{"data": deployment}
	}
}

func RollbackDeployment(request contracts.HttpRequest, guard contracts.Guard) any {
	var form requests.RollbackDeploymentRequest
	if err := request.Parse(&form); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	deployment := models.Deployments().FindOrFail(form.Id)
	project := models.Projects().FindOrFail(deployment.ProjectId)

	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	if deployment.Status != models.StatusFinished {
		return contracts.Fields{"msg": "不允许撤回该部署记录"}
	}

	if outputs, err := usecase.RollbackDeployment(
		project,
		deployment,
		form.Commands,
		form.BeforeRelease,
		form.AfterRelease,
	); err != nil {
		return contracts.Fields{"msg": err.Error(), "data": outputs}
	} else {
		return contracts.Fields{"data": outputs}
	}
}

func RunDeployment(request contracts.HttpRequest, guard contracts.Guard) any {
	deployment := models.Deployments().FindOrFail(request.Get("id"))
	project := models.Projects().FindOrFail(deployment.ProjectId)
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	go usecase.GoDeployment(deployment, models.Commands().Where("project_id", deployment.ProjectId).Get())

	return contracts.Fields{
		"data": "ok",
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
