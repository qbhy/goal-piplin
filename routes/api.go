package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/http/controllers"
	"github.com/qbhy/goal-piplin/app/http/controllers/manage"
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

	authRouter.Get("/project/list", controllers.GetProjects)
	authRouter.Get("/project/detail", controllers.GetProject)
	authRouter.Post("/project/create", controllers.CreateProject)
	authRouter.Post("/project/update", controllers.UpdateProject)

	authRouter.Get("/deployment/list", controllers.GetDeployments)
	authRouter.Get("/deployment/detail", controllers.GetDeploymentDetail)
	authRouter.Post("/deployment/create", controllers.CreateDeployment)

	authRouter.Get("/environment/list", controllers.GetEnvironments)
	authRouter.Post("/environment/create", controllers.CreateEnvironment)
	authRouter.Post("/environment/update", controllers.UpdateEnvironment)
	authRouter.Post("/environment/delete", controllers.DeleteEnvironment)

	authRouter.Get("/config/list", controllers.GetConfigs)
	authRouter.Post("/config/create", controllers.CreateConfig)
	authRouter.Post("/config/update", controllers.UpdateConfig)
	authRouter.Post("/config/delete", controllers.DeleteConfig)

	authRouter.Get("/share/list", controllers.GetShares)
	authRouter.Post("/share/create", controllers.CreateShare)
	authRouter.Post("/share/update", controllers.UpdateShare)
	authRouter.Post("/share/delete", controllers.DeleteShare)

	authRouter.Get("/command/list", controllers.GetCommands)
	authRouter.Post("/command/create", controllers.CreateCommand)
	authRouter.Post("/command/update", controllers.UpdateCommand)
	authRouter.Post("/command/delete", controllers.DeleteCommand)

	authRouter.Get("/cabinet/list", manage.GetCabinets)
	authRouter.Post("/cabinet/create", manage.CreateCabinet)
	authRouter.Post("/cabinet/update", manage.UpdateCabinet)
	authRouter.Post("/cabinet/delete", manage.DeleteCabinet)

	authRouter.Get("/key/list", manage.GetKeys)
	authRouter.Post("/key/create", manage.CreateKey)
	authRouter.Post("/key/delete", manage.DeleteKeys)
	authRouter.Post("/key/update", manage.UpdateKey)

	authRouter.Get("/group/list", manage.GetGroups)
	authRouter.Post("/group/create", manage.CreateGroup)
	authRouter.Post("/group/update", manage.UpdateGroup)
	authRouter.Post("/group/delete", manage.DeleteGroups)

}
