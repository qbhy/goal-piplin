package manage

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/usecase"
)

func GetGroups(request contracts.HttpRequest) any {
	return contracts.Fields{
		"data": models.Groups().
			When(request.GetString("name") != "", func(q contracts.QueryBuilder[models.Group]) contracts.Query[models.Group] {
				return q.Where("name", "like", "%"+request.GetString("name")+"%")
			}).
			Get().ToArray(),
	}
}

func CreateGroup(request contracts.HttpRequest, guard contracts.Guard) any {
	return contracts.Fields{
		"data": models.Groups().Create(contracts.Fields{
			"name":       request.GetString("name"),
			"creator_id": guard.User().GetId(),
		}),
	}
}

func DeleteGroups(request contracts.HttpRequest) any {
	err := usecase.DeleteGroups(request.Get("id"))

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}

func UpdateGroup(request contracts.HttpRequest) any {
	err := usecase.UpdateGroup(request.Get("id"), request.Fields())

	if err != nil {
		return contracts.Fields{"msg": err.Error()}
	}

	return contracts.Fields{"data": nil}
}
