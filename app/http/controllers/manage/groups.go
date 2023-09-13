package manage

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func GetGroups() any {
	return models.Groups().Get().ToArray()
}

func CreateGroup(request contracts.HttpRequest) any {
	return models.Groups().Get()
}
