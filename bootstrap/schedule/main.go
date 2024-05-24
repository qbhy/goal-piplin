package main

import (
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal-piplin/bootstrap/core"
)

func main() {

	app := core.Application()

	if errors := app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Fatal("goal 异常!")
	} else {
		logs.Default().Info("goal 已关闭")
	}
}
