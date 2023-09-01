package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
	"time"
)

var CommandClass = class.Make[Command]()

func CommandQuery() *table.Table[Command] {
	return table.Class(CommandClass, "commands")
}

type Command struct {
	Id              string    `json:"id"`
	ProjectId       int       `json:"project_id"`       // 项目ID
	Step            string    `json:"step"`             // 步骤
	Sort            int       `json:"sort"`             // 排序
	Name            string    `json:"name"`             // 名称
	User            string    `json:"user"`             // 运行用户
	Script          string    `json:"script"`           // shell 脚本
	Environments    string    `json:"environments"`     // 环境
	Optional        bool      `json:"optional"`         // 是否可选
	DefaultSelected bool      `json:"default_selected"` // 是否默认选中
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
