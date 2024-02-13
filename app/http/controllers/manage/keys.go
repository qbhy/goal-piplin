package manage

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/example/app/usecase"
	"time"
)

func GetKeys() any {
	groups := models.Keys().Get().ToArray()
	return groups
}

func CreateKey(request contracts.HttpRequest, guard contracts.Guard) any {
	privateKey, publicKey, err := usecase.GenerateRSAKey()
	if err != nil {
		panic(err)
	}
	return models.Keys().Create(contracts.Fields{
		"name":        request.GetString("name"),
		"public_key":  string(publicKey),
		"private_key": string(privateKey),
		"created_at":  time.Now(),
	})
}
