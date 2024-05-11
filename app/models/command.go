package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var CommandClass = class.Make[Command]()

func Commands() *table.Table[Command] {
	return table.Class(CommandClass, "commands")
}

type Command struct {
	table.Model[Command] `json:"-"`

	Id              int    `json:"id"`
	Name            string `json:"name"`             // 名称
	ProjectId       int    `json:"project_id"`       // 项目ID
	Step            string `json:"step"`             // 步骤
	Sort            int    `json:"sort"`             // 排序
	User            string `json:"user"`             // 运行用户
	Script          string `json:"script"`           // shell 脚本
	Environments    []int  `json:"environments"`     // 环境
	Optional        bool   `json:"optional"`         // 是否可选
	DefaultSelected bool   `json:"default_selected"` // 是否默认选中
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
