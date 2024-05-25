package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
	"github.com/goal-web/http/sse"
	"github.com/goal-web/http/websocket"
	"github.com/goal-web/session"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/bootstrap/core"
	"github.com/qbhy/goal-piplin/routes"
)

func main() {

	app := core.Application()

	app.RegisterServices(
		http.NewService(routes.Api, routes.WebSocket, routes.Sse),
		session.NewService(),
		sse.NewService(),
		websocket.NewService(),
		//signal.NewService(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher, console3 contracts.Console, input contracts.ConsoleInput) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		console3.Run(input)
	})
}
