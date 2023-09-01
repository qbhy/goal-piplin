package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
	"time"
)

var ProjectClass = class.Make[Project]()

func ProjectQuery() *table.Table[Project] {
	return table.Class(ProjectClass, "projects")
}

type Project struct {
	Id            string    `json:"id"`
	Uuid          string    `json:"uuid"`
	Name          string    `json:"name"`
	CreatorId     int       `json:"creator_id"`
	PublicKey     string    `json:"public_key"`
	PrivateKey    string    `json:"private_key"`
	RepoAddress   string    `json:"repo_address"`
	ProjectPath   string    `json:"project_path"`
	DefaultBranch string    `json:"default_branch"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
