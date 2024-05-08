package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
	"time"
)

func init() {
	configs["http"] = func(env contracts.Env) any {
		return http.Config{
			Host:         env.GetString("http.host"),
			Port:         env.GetString("http.port"),
			SseHeartBeat: 15 * time.Second,
		}
	}
}
