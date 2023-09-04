package usecase

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func Login(user *models.User, guard contracts.Guard) contracts.Fields {
	return contracts.Fields{
		"user":  user,
		"token": guard.Login(user),
	}
}
