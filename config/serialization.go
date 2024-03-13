package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/serialization"
	"github.com/qbhy/goal-piplin/app/jobs"
)

func init() {
	configs["serialization"] = func(env contracts.Env) any {
		return serialization.Config{
			Default: "json", // 支持：json、gob、xml。
			Class: []contracts.Class[any]{ // 需要序列化的类
				jobs.DemoClass,
			},
		}
	}
}
