package requests

import (
	"github.com/goal-web/contracts"
)

type CreateUserRequest struct {
	contracts.HttpRequest `di:""` // 加入 di 标记表示需要注入
}

func (req CreateUserRequest) Rules() contracts.Fields {
	return contracts.Fields{
		"username": "required",
		"role":     "required",
		"password": "required",
	}
}
