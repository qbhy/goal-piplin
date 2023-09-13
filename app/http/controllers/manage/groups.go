package manage

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func GetGroups() any {
	groups := models.Groups().Get().ToArray()
	groups = append(groups, models.Group{Id: 0, Name: "未分组"})
	return groups
}

func CreateGroup(request contracts.HttpRequest, guard contracts.Guard) any {
	return models.Groups().Create(contracts.Fields{
		"name":       request.GetString("name"),
		"creator_id": guard.User().GetId(),
	})
}
