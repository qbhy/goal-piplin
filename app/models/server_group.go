package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var CabinetClass = class.Make[Cabinet]()

func Cabinets() *table.Table[Cabinet] {
	return table.Class(CabinetClass, "cabinets")
}

type CabinetSettings struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	Enabled bool   `json:"enabled"`
}

type Cabinet struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`     // 名称
	Settings  []CabinetSettings `json:"settings"` // 配置
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
