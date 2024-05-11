package middlewares

import (
	"errors"
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
)

func Admin(request contracts.HttpRequest, next contracts.Pipe, guard contracts.Guard) any {
	if guard.Guest() {
		panic(auth.Exception{Err: errors.New("guard authentication failed")})
	}

	user := guard.User().(*models.User)

	if user.Role == models.UserRoleAdmin {
		return next(request)
	}

	return contracts.Fields{"msg": "您不是管理员"}
}
