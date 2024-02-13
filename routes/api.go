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
	authRouter.Get("/project/{id}", controllers.GetProject)
	authRouter.Post("/project", controllers.CreateProject)

	authRouter.Get("/deployments", controllers.GetDeployments)
	authRouter.Get("/environments", controllers.GetEnvironments)
	authRouter.Post("/environment", controllers.CreateEnvironment)
	authRouter.Get("/cabinet/list", controllers.GetCabinets)
	authRouter.Post("/cabinet/create", controllers.CreateCabinet)
	authRouter.Post("/cabinet/update", controllers.UpdateCabinet)
	authRouter.Post("/cabinet/delete", controllers.DeleteCabinet)

	manageRouter := authRouter.Group("/manage", middlewares.Manage)
	{
		manageRouter.Get("/groups", manage.GetGroups)
		manageRouter.Post("/group", manage.CreateGroup)
		manageRouter.Get("/keys", manage.GetKeys)
		manageRouter.Post("/key", manage.CreateKey)

	}

}
