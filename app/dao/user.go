package dao

import "github.com/goal-web/example/app/models"

func FindUser(id any) *models.User {
	return models.UserQuery().Find(id)
}
