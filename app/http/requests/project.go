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
