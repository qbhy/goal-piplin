package manage

import (
	"github.com/goal-web/contracts"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
	"github.com/qbhy/goal-piplin/app/utils"
)

func GetKeys(request contracts.HttpRequest) any {
	list, total := models.Keys().
		OrderByDesc("id").
		Paginate(20, request.Int64Optional("page", 1))
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
			"name":        request.GetString("name"),
			"public_key":  publicKey,
			"private_key": privateKey,
			"created_at":  carbon.Now().ToDateTimeString(),
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
