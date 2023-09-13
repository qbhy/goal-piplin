package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var GroupClass = class.Make[Group]()

func Groups() *table.Table[Group] {
	return table.Class(GroupClass, "`groups`")
}

type Group struct {
	Id        string `json:"id"`
	Name      string `json:"name"`    // 名称
	CreatorId int    `json:"creator"` // 创建者
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
