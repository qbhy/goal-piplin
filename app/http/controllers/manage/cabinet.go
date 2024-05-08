package manage

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetCabinets(request contracts.HttpRequest, guard contracts.Guard) any {
	user := guard.User().(models.User)
	list, total := models.Cabinets().
		OrderByDesc("id").
		When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.Cabinet]) contracts.Query[models.Cabinet] {
			return q.Where("name", "like", "%"+request.GetString("name")+"%")
		}).
		When(user.Role != "admin", func(q contracts.QueryBuilder[models.Cabinet]) contracts.Query[models.Cabinet] {
			return q.Where("creator_id", user.Id)
		}).
		Paginate(request.Int64Optional("pageSize", 10), request.Int64Optional("current", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateCabinet(request contracts.HttpRequest, guard contracts.Guard) any {
	cabinet, err := usecase.CreateCabinet(guard.GetId(), request.GetString("name"), request.Get("settings"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": cabinet}
}

func UpdateCabinet(request contracts.HttpRequest) any {
	err := usecase.UpdateCabinet(request.Get("id"), request.GetString("name"), request.Get("settings"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func DeleteCabinet(request contracts.HttpRequest) any {
	err := usecase.DeleteCabinet(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
