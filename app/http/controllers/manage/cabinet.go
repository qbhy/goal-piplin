package manage

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetCabinets(request contracts.HttpRequest) any {
	list, total := models.Cabinets().
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
	return contracts.Fields{
		"total": total,
		"data":  list.ToArray(),
	}
}

func CreateCabinet(request contracts.HttpRequest) any {
	cabinet, err := usecase.CreateCabinet(request.GetString("name"), request.Get("settings"))

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
