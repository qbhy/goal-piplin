package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal-piplin/app/models"
	utils2 "github.com/qbhy/goal-piplin/app/utils"
)

func CreateKey(name string) (models.Key, error) {
	var key models.Key
	privateKey, publicKey, err := utils2.GenerateRSAKeys()
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

func DeleteKeys(id any) error {
	if models.Projects().WhereIn("key_id", id).Count() > 0 {
		return errors.New("不能删除该密钥，因为有项目正在使用此公钥。")
	}
	_, err := models.Keys().WhereIn("id", id).DeleteE()
	return err
}

func UpdateKey(id any, fields contracts.Fields) error {
	if models.Keys().Where("id", "!=", id).Where("name", fields["name"]).Count() > 0 {
		return errors.New("密钥名称已存在")
	}

	_, err := models.Keys().Where("id", id).UpdateE(utils.OnlyFields(fields, "name", "public_key", "private_key"))

	return err
}
