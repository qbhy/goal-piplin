package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/migration"
)

func init() {
	configs["migration"] = func(env contracts.Env) any {
		return &migration.Config{
			Dir: "database/migrations",
		}
	}
}
