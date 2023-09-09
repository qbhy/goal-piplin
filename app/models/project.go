package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var ProjectClass = class.Make[Project]()

func Projects() *table.Table[Project] {
	return table.Class(ProjectClass, "projects")
}

type ProjectEnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Callback struct {
	Name    string   `json:"name"`
	Webhook string   `json:"webhook"`
	Events  []string `json:"events"`
	Enabled bool     `json:"enabled"`
}

type ProjectSettings struct {
	EnvVars   []ProjectEnvVar `json:"vars"`
	Branches  []string        `json:"branches"`
	Callbacks []Callback      `json:"callbacks"`
}

type Project struct {
	Id            string          `json:"id"`
	Uuid          string          `json:"uuid"`
	Settings      ProjectSettings `json:"settings"`
	Name          string          `json:"name"`
	CreatorId     int             `json:"creator_id"`
	KeyId         int             `json:"key_id"`
	RepoAddress   string          `json:"repo_address"`
	ProjectPath   string          `json:"project_path"`
	DefaultBranch string          `json:"default_branch"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
}
