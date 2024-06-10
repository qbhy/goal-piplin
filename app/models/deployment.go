package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

const (
	BeforeClone   = "before_clone"
	AfterClone    = "after_clone"
	BeforePrepare = "before_prepare"
	AfterPrepare  = "after_prepare"
	BeforeRelease = "before_release"
	AfterRelease  = "after_release"

	Init    = "init"    // 创建目录
	Clone   = "clone"   // 克隆代码
	Prepare = "prepare" // 准备配置文件、共享文件等
	Release = "release" // 切换版本

	StatusWaiting  = "waiting"
	StatusRunning  = "running"
	StatusFailed   = "failed"
	StatusFinished = "finished"
)

var DeploymentClass = class.Make[Deployment]()

func Deployments() *table.Table[Deployment] {
	return table.Class(DeploymentClass, "deployments")
}

type Deployment struct {
	table.Model[Deployment] `json:"-"`

	Id           string          `json:"id"`
	ProjectId    int             `json:"project_id"` // 项目ID
	Version      string          `json:"version"`    // 部署版本
	Comment      string          `json:"comment"`    // 说明
	Commit       string          `json:"commit"`     // 提交 hash
	Status       string          `json:"status"`     // 状态
	Params       map[string]bool `json:"params"`     // {step: bool}
	Results      []CommandResult `json:"results"`
	Environments []int           `json:"environments"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

type CommandResult struct {
	Step          string                   `json:"step"`
	Servers       map[string]CommandOutput `json:"servers"` // { ip : CommandOutput}
	Command       int                      `json:"command,omitempty"`
	Name          string                   `json:"name,omitempty"`
	TimeConsuming int                      `json:"time_consuming,omitempty"` // 用时，单位秒
}

type CommandOutput struct {
	Server
	Outputs string `json:"outputs,omitempty"`
	Status  string `json:"status"`
	Time    int    `json:"time,omitempty"`
}
