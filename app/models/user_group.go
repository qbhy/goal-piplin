package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var UserGroupClass = class.Make[UserGroup]()

func UserGroups() *table.Table[UserGroup] {
	return table.Class(UserGroupClass, "user_groups")
}

type UserGroup struct {
	Id        string `json:"id"`
	GroupId   int    `json:"group_id"` // 项目ID
	UserId    int    `json:"user_id"`  // 用户ID
	Status    string `json:"status"`   // 状态：inviting、joined、rejected
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
