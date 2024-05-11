package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var ShareFileClass = class.Make[ShareFile]()

func ShareFiles() *table.Table[ShareFile] {
	return table.Class(ShareFileClass, "share_files")
}

type ShareFile struct {
	table.Model[ShareFile] `json:"-"`

	Id        string `json:"id"`
	ProjectId int    `json:"project_id"` // 项目ID
	Name      string `json:"name"`       // 名称
	Path      string `json:"path"`       // 文件路径
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
