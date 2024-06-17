package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/querybuilder"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/validation"
	"github.com/qbhy/goal-piplin/app/http/requests"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetProjects(request contracts.HttpRequest, guard contracts.Guard) any {
	user := guard.User().(*models.User)

	list, total := models.Projects().
		OrderByDesc("id").
		When(user.Role != "admin", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {

			return q.WhereFunc(func(q contracts.Query[models.Project]) {
				q.Where("creator_id", user.Id).
					OrWhereExists(func() contracts.Query[models.Project] {
						return querybuilder.New[models.Project]("user_projects").
							Where("user_id", user.Id).
							Where("status", models.InviteStatusJoined).
							Where("projects.id", querybuilder.Expression("user_projects.project_id"))
					}).
					OrWhereExists(func() contracts.Query[models.Project] {
						return querybuilder.New[models.Project]("user_groups").
							Where("user_id", user.Id).
							Where("status", models.InviteStatusJoined).
							Where("projects.group_id", querybuilder.Expression("user_groups.group_id"))
					})
			})
		}).
		When(request.GetString("group_id") != "", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {
			return q.Where("group_id", request.GetString("group_id"))
		}).
		When(request.GetString("name") != "", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(request.GetString("repo_address") != "", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {
			return q.Where("repo_address", "like", "%"+request.GetString("repo_address")+"%")
		}).
		When(request.GetString("project_path") != "", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {
			return q.Where("project_path", "like", "%"+request.GetString("project_path")+"%")
		}).
		When(request.GetString("default_branch") != "", func(q contracts.Query[models.Project]) contracts.Query[models.Project] {
			return q.Where("default_branch", "like", "%"+request.GetString("default_branch")+"%")
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func GetProject(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.GetString("id"))
	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	return contracts.Fields{
		"data": usecase.GetProjectDetail(project),
	}
}

func DeleteProject(request contracts.HttpRequest, guard contracts.Guard) any {
	project := models.Projects().FindOrFail(request.Get("id"))

	if !usecase.HasProjectPermission(project, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	err := usecase.DeleteProject(project)
	if err != nil {
		return contracts.Fields{"msg": "删除失败：" + err.Error()}
	}
	return contracts.Fields{"successful": true}
}

func CreateProject(request requests.ProjectRequest, guard contracts.Guard) any {
	validation.VerifyForm(request)
	fields := request.Fields()

	if groupId := utils.ToInt(fields["group_id"], 0); groupId > 0 && !usecase.HasGroupPermission(models.Groups().FindOrFail(groupId), guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该分组的权限"}
	}

	project, err := usecase.CreateProject(guard.GetId(), fields)

	if err != nil {
		return contracts.Fields{"msg": "创建项目失败：" + err.Error()}
	}

	return contracts.Fields{
		"data": project,
	}
}

func CopyProject(request contracts.HttpRequest, guard contracts.Guard) any {
	var form requests.CopyProjectForm

	if err := request.Parse(&form); err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	targetProject := models.Projects().FindOrFail(form.TargetProject)

	if !usecase.HasProjectPermission(targetProject, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}

	if form.GroupId > 0 && !usecase.HasGroupPermission(models.Groups().FindOrFail(form.GroupId), guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该分组的权限"}
	}

	project, err := usecase.CopyProject(targetProject, contracts.Fields{
		"name":           form.Name,
		"key_id":         form.KeyId,
		"repo_address":   form.RepoAddress,
		"default_branch": form.DefaultBranch,
		"group_id":       form.GroupId,
		"creator_id":     guard.GetId(),
	})

	if err != nil {
		return contracts.Fields{"msg": "复制项目：" + err.Error()}
	}

	return contracts.Fields{
		"data": project,
	}
}

func UpdateProject(request requests.ProjectRequest, guard contracts.Guard) any {
	validation.VerifyForm(request)
	fields := request.Fields()
	targetProject := models.Projects().FindOrFail(request.GetString("id"))
	if !usecase.HasProjectPermission(targetProject, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该项目的权限"}
	}
	project, err := usecase.UpdateProject(request.GetInt("id"), fields)
	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}
	return contracts.Fields{"data": project}
}
