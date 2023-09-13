package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/controllers"
	"github.com/goal-web/example/app/http/controllers/manage"
	"github.com/goal-web/example/app/http/middlewares"
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

	authRouter.Get("/projects", controllers.GetProjects)
	authRouter.Post("/project", controllers.CreateProject)

	manageRouter := authRouter.Group("/manage", middlewares.Manage)
	{
		manageRouter.Get("/groups", manage.GetGroups)
		manageRouter.Post("/group", manage.CreateGroup)
	}

}
