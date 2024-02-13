package usecase

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func CreateKey(name string) (models.Key, error) {
	var key models.Key
	publicKey, privateKey, err := GenerateRSAKey()
	if err != nil {
		return key, err
	}
	key = models.Keys().Create(contracts.Fields{
		"name":        name,
		"public_key":  publicKey,
		"private_key": privateKey,
	})
	return key, nil
}
