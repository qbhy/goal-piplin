package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var DeploymentClass = class.Make[Deployment]()

func Deployments() *table.Table[Deployment] {
	return table.Class(DeploymentClass, "deployments")
}

type Deployment struct {
	Id           string `json:"id"`
	ProjectId    int    `json:"project_id"` // 项目ID
	Version      string `json:"version"`    // 部署版本
	Comment      string `json:"comment"`    // 说明
	Environments string `json:"environments"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
