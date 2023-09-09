package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var ServerGroupClass = class.Make[ServerGroup]()

func ServerGroups() *table.Table[ServerGroup] {
	return table.Class(ServerGroupClass, "server_groups")
}

type ServerGroupSettings struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	Enabled bool   `json:"enabled"`
}

type ServerGroup struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`     // 名称
	Settings  ServerGroupSettings `json:"settings"` // 配置
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
}
