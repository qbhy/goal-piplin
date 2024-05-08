package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var CabinetClass = class.Make[Cabinet]()

func Cabinets() *table.Table[Cabinet] {
	return table.Class(CabinetClass, "cabinets")
}

type Cabinet struct {
	Id        string   `json:"id"`
	CreatorId int      `json:"creator_id"`
	Name      string   `json:"name"`     // 名称
	Settings  []Server `json:"settings"` // 配置
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
