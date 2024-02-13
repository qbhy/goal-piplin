package tests

import (
	"github.com/goal-web/application"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	console2 "github.com/goal-web/example/app/console"
	config2 "github.com/goal-web/example/config"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/hashing"
	"github.com/goal-web/redis"
	"github.com/goal-web/session"
	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/logs"
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
			config.NewToml(config.File("config.toml")),
			config2.GetConfigProviders(),
		),
		console.NewService(console2.NewKernel),
		hashing.NewService(),
		encryption.NewService(),
		filesystem.NewService(),
		events.NewService(),
		redis.NewService(),
		bloomfilter.NewService(),
		cache.NewService(),
		session.NewService(),
		auth.NewService(),
		email.NewService(),
		database.NewService(),
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
