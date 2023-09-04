package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/controllers"
)

func Api(router contracts.HttpRouter) {
	api := router.Group("/api")
	api.Post("/login", controllers.Login)

	api.Get("/", controllers.HelloWorld)

	api.Post("/queue", controllers.DemoJob)

	api.Get("/micro", controllers.RpcService)
	//router.Get("/", controllers.HelloWorld, ratelimiter.Middleware(100))

	authRouter := api.Group("", auth.Guard("jwt"))
	authRouter.Get("/myself", controllers.GetCurrentUser)

	router.Post("/mail", controllers.SendEmail)
}
