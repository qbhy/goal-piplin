package main

import (
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/hashing"
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/goal-web/supports/logs"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/console"
	"github.com/qbhy/goal-piplin/app/exceptions"
	"github.com/qbhy/goal-piplin/app/providers"
	config2 "github.com/qbhy/goal-piplin/config"
)

func main() {
	env := config.NewToml(config.File("env.toml"))
	app := application.Singleton(env.GetBool("app.debug"))

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(env, config2.GetConfigProviders()),
		events.NewService(),
		providers.NewEvents(),
		hashing.NewService(),
		encryption.NewService(),
		filesystem.NewService(),
		serialization.NewService(),
		redis.NewService(),
		cache.NewService(),
		bloomfilter.NewService(),
		auth.NewService(),
		ratelimiter.NewService(),
		console.NewService(),
		database.NewService(),
		queue.NewService(false),
		email.NewService(),
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
	})
	if errors := app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Fatal("goal 异常!")
	} else {
		logs.Default().Info("goal 已关闭")
	}
}
