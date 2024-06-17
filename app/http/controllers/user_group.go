package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/table"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetUserGroups(request contracts.HttpRequest, guard contracts.Guard) any {
	user := guard.User().(*models.User)
	list, total := table.ArrayQuery("user_groups").
		OrderByDesc("user_groups.id").
		Select("group_id", "`groups`.name as group_name", "user_groups.created_at", "user_groups.id", "status").
		LeftJoin("`groups`", "`groups`.id", "=", "user_groups.group_id").
		When(request.GetString("group_name") != "", func(q contracts.Query[contracts.Fields]) contracts.Query[contracts.Fields] {
			return q.Where("groups.name", "like", "%"+request.GetString("group_name")+"%")
		}).
		When(request.GetString("status") != "", func(q contracts.Query[contracts.Fields]) contracts.Query[contracts.Fields] {
			return q.Where("status", request.GetString("status"))
		}).
		Where("user_id", user.Id).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateUserGroup(request contracts.HttpRequest, guard contracts.Guard) any {
	userId := request.GetInt("user_id")
	group := models.Groups().FindOrFail(request.GetInt("group_id"))

	if !usecase.HasGroupPermission(group, guard.User().(*models.User)) {
		return contracts.Fields{"msg": "没有该分组的权限"}
	}

	data, err := usecase.CreateUserGroup(group.Id, userId)

	if request.GetString("user_id") == guard.GetId() {
		return contracts.Fields{"msg": "不可以邀请自己"}
	}
	if group.CreatorId == userId {
		return contracts.Fields{"msg": "不可以邀请创建者"}
	}

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}
	return contracts.Fields{"data": data}
}

func DeleteUserGroups(request contracts.HttpRequest) any {
	err := usecase.DeleteUserGroup(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func UpdateUserGroup(request contracts.HttpRequest, guard contracts.Guard) any {
	userGroup := models.UserGroups().
		Where("group_id", request.Get("group_id")).
		Where("user_id", guard.GetId()).FirstOrFail()

	err := usecase.UpdateUserGroup(userGroup, request.GetString("status"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
