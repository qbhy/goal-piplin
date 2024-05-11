package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var UserProjectClass = class.Make[UserProject]()

const (
	InviteStatusWaiting  = "inviting"
	InviteStatusJoined   = "joined"
	InviteStatusRejected = "rejected"
)

func UserProjects() *table.Table[UserProject] {
	return table.Class(UserProjectClass, "user_projects")
}

type UserProject struct {
	table.Model[UserProject] `json:"-"`

	Id        string `json:"id"`
	ProjectId int    `json:"project_id"` // 项目ID
	UserId    int    `json:"user_id"`    // 用户ID
	Status    string `json:"status"`     // 状态：inviting、joined、rejected
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
