package usecase

import (
	"errors"
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/models"
)

func CreateUser(name, password, role string) (*models.User, error) {
	if models.Users().Where("username", name).Count() > 0 {
		return nil, errors.New("用户名已存在")
	}

	user, err := models.Users().CreateE(contracts.Fields{
		"username":   name,
		"nickname":   name,
		"avatar":     "",
		"role":       role,
		"password":   application.Get("hashing").(contracts.Hasher).Make(password, nil),
		"created_at": carbon.Now().ToDateTimeString(),
	})
	return user, err
}

func DeleteUsers(id any) error {
	if models.Projects().Where("creator_id", id).Count() > 0 {
		return errors.New("不能删除该用户，请先删除该用户创建的项目。")
	}
	_, err := models.Users().Where("id", id).DeleteE()
	return err
}

func UpdateUser(id any, fields contracts.Fields) error {
	_, err := models.Users().Where("id", id).UpdateE(utils.OnlyFields(fields, "nickname", "password", "avatar"))

	return err
}
