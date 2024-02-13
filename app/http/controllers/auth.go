package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/http/requests"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/example/app/usecase"
	"github.com/goal-web/validation"
)

func Login(guard contracts.Guard, request requests.LoginRequest, hash contracts.Hasher) any {
	// 验证不通过将抛异常，如希望自己处理错误，请使用 Form 方法
	validation.VerifyForm(request)

	//  这是伪代码
	var user, e = models.Users().FirstWhereE("username", request.GetString("username"))
	if e != nil {
		return contracts.Fields{"msg": "账号或者密码不正确"}
	}

	if hash.Check(request.GetString("password"), user.Password, nil) {
		return usecase.Login(user, guard)
	}

	return contracts.Fields{"msg": "账号或者密码不正确"}
}

func GetCurrentUser(guard contracts.Guard) any {
	return contracts.Fields{
		"user": guard.User(),
	}
}
