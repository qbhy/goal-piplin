package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/bootstrap/core"
)

func main() {
	app := core.Application()

	app.RegisterServices()

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
	})

	app.Call(func(console3 contracts.Console, input contracts.ConsoleInput) {
		console3.Run(input)
	})
}
