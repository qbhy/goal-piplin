package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
	"time"
)

var ConfigFileClass = class.Make[ConfigFile]()

func ConfigFiles() *table.Table[ConfigFile] {
	return table.Class(ConfigFileClass, "config_files")
}

type ConfigFile struct {
	Id           string    `json:"id"`
	ProjectId    int       `json:"project_id"`   // 项目ID
	Name         string    `json:"name"`         // 名称
	Path         string    `json:"path"`         // 文件路径
	Content      string    `json:"content"`      // 内容
	Environments string    `json:"environments"` // 关联环境
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
