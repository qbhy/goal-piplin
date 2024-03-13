package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var ProjectEnvironmentClass = class.Make[ProjectEnvironment]()

func ProjectEnvironments() *table.Table[ProjectEnvironment] {
	return table.Class(ProjectEnvironmentClass, "project_environments")
}

type Server struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Disabled    bool   `json:"disabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Environment int    `json:"-"`
}

type EnvironmentSettings struct {
	Servers         []Server `json:"servers"`          // 服务器列表
	Cabinets        []string `json:"cabinets"`         // 服务器租
	DefaultSelected bool     `json:"default_selected"` // 默认选中
	LinkageDeploy   string   `json:"linkage_deploy"`   // 联动部署环境
}

type ProjectEnvironment struct {
	Id        int                 `json:"id"`
	ProjectId int                 `json:"project_id"`
	Name      string              `json:"name"`
	Settings  EnvironmentSettings `json:"settings"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
}
