package manage

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/validation"
	"github.com/qbhy/goal-piplin/app/http/requests"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetUsers(request contracts.HttpRequest) any {
	list, total := models.Users().
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.Query[models.User]) contracts.Query[models.User] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateUser(request requests.CreateUserRequest) any {
	validation.VerifyForm(request)

	newUser, err := usecase.CreateUser(request.GetString("username"), request.GetString("password"), request.GetString("role"))
	if err != nil {
		return contracts.Fields{"msg": "创建用户失败：" + err.Error()}
	}
	return contracts.Fields{
		"data": newUser,
	}
}

func DeleteUsers(request contracts.HttpRequest) any {
	err := usecase.DeleteUsers(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func UpdateUser(request contracts.HttpRequest, hash contracts.Hasher) any {
	fields := request.Fields()
	if fields["password"] != nil {
		fields["password"] = hash.Make(utils.ToString(fields["password"], ""), nil)
	}
	err := usecase.UpdateUser(request.Get("id"), fields)

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
