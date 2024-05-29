package requests

import (
	"github.com/goal-web/contracts"
)

type ProjectRequest struct {
	contracts.HttpRequest `di:""` // 加入 di 标记表示需要注入
}

func (req ProjectRequest) Rules() contracts.Fields {
	return contracts.Fields{
		//"group_id":       "required",
		//"key_id":         "required",
		"name":           "required",
		"repo_address":   "required",
		"project_path":   "required",
		"default_branch": "required",
	}
}
func (req ProjectRequest) Fields() contracts.Fields {
	return contracts.Fields{
		"group_id":       req.GetInt("group_id"),
		"key_id":         req.GetInt("key_id"),
		"name":           req.GetString("name"),
		"repo_address":   req.GetString("repo_address"),
		"project_path":   req.GetString("project_path"),
		"default_branch": req.GetString("default_branch"),
	}
}
