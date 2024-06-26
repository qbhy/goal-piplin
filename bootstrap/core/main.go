package core

import (
	"flag"
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console/scheduling"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/hashing"
	"github.com/goal-web/migration"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/console"
	"github.com/qbhy/goal-piplin/app/exceptions"
	"github.com/qbhy/goal-piplin/app/providers"
	config2 "github.com/qbhy/goal-piplin/config"
)

var envPath = flag.String("env", "env.toml", "指定 env")

func Application() contracts.Application {
	flag.Parse()

	env := config.NewToml(config.File(*envPath))
	app := application.Singleton(env.GetBool("app.debug"))

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(env, config2.GetConfigProviders()),
		hashing.NewService(),
		encryption.NewService(),
		filesystem.NewService(),
		serialization.NewService(),
		events.NewService(),
		providers.NewEvents(),
		redis.NewService(),
		cache.NewService(),
		bloomfilter.NewService(),
		auth.NewService(),
		ratelimiter.NewService(),
		console.NewService(),
		scheduling.NewService(),
		database.NewService(),
		migration.NewService(),
		email.NewService(),
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
	})

	return app
}
