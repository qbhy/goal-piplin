package manage

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
	"github.com/qbhy/goal-piplin/app/utils"
)

func GetKeys(request contracts.HttpRequest, guard contracts.Guard) any {
	user := guard.User().(*models.User)
	list, total := models.Keys().
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.Key]) contracts.Query[models.Key] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(user.Role != "admin", func(q contracts.QueryBuilder[models.Key]) contracts.Query[models.Key] {
			return q.Where("creator_id", user.Id)
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateKey(request contracts.HttpRequest, guard contracts.Guard) any {
	privateKey, publicKey, err := utils.GenerateRSAKeys()
	if err != nil {
		panic(err)
	}
	return contracts.Fields{
		"data": models.Keys().Create(contracts.Fields{
			"creator_id":  guard.GetId(),
			"name":        request.GetString("name"),
			"public_key":  publicKey,
			"private_key": privateKey,
		}),
	}
}

func DeleteKeys(request contracts.HttpRequest) any {
	err := usecase.DeleteKeys(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func UpdateKey(request contracts.HttpRequest) any {
	err := usecase.UpdateKey(request.Get("id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
