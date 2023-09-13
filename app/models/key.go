package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var KeyClass = class.Make[Key]()

func Keys() *table.Table[Key] {
	return table.Class(KeyClass, "`keys`")
}

type Key struct {
	Id         string `json:"id"`
	Name       string `json:"name"`        // 名称
	PublicKey  string `json:"public_key"`  // 公钥
	PrivateKey string `json:"private_key"` // 私钥
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
