package tests

import (
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
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/serialization"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal-piplin/app/console"
	"github.com/qbhy/goal-piplin/app/providers"
	config2 "github.com/qbhy/goal-piplin/config"
	"sync"
)

func initApp(path ...string) contracts.Application {
	app := application.Singleton()

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.DefaultExceptionHandler{}
	})

	wg := sync.WaitGroup{}

	app.RegisterServices(
		config.NewService(
			config.NewToml(config.File("env.toml")),
			config2.GetConfigProviders(),
		),
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
		queue.NewService(true),
		email.NewService(),
		//&http.serviceProvider{RouteCollectors: []any{
		//	// 路由收集器
		//	routes.V1Routes,
		//}},
	)

	wg.Add(1)
	go func() {
		wg.Done()
		if errors := app.Start(); len(errors) > 0 {
			logs.WithField("errors", errors).Fatal("goal 启动异常!")
		}
	}()
	wg.Wait()
	return app
}
