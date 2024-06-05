package middlewares

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/qbhy/goal-piplin/app/models"
)

func Admin(request contracts.HttpRequest, next contracts.Pipe, guard contracts.Guard) any {
	if guard.Guest() {
		panic(auth.Exception{Exception: exceptions.New("guard authentication failed")})
	}

	user := guard.User().(*models.User)

	if user.Role == models.UserRoleAdmin {
		return next(request)
	}

	return contracts.Fields{"msg": "您不是管理员"}
}
