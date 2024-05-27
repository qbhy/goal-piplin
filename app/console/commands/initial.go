package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal-piplin/app/models"
	"github.com/qbhy/goal-piplin/app/utils"
)

func NewInitial(app contracts.Application) contracts.Command {
	return &Initial{
		Command: commands.Base("init", "初始化 goal-piplin"),
		hash:    app.Get("hashing").(contracts.Hasher),
	}
}

type Initial struct {
	commands.Command
	hash contracts.Hasher
}

func (cmd Initial) Handle() any {
	username := "piplin"
	password := "password"
	user := models.Users().Where("username", username).First()
	var err error
	if user != nil {
		logs.Default().Info("piplin 用户已存在")
	} else {
		user, err = models.Users().CreateE(contracts.Fields{
			"username": username,
			"nickname": username,
			"avatar":   "",
			"role":     models.UserRoleAdmin,
			"password": cmd.hash.Make(password, nil),
		})
		if err != nil {
			panic(err)
		}
		logs.Default().Info(fmt.Sprintf("已创建用户 %s 密码为 %s", username, password))
	}

	if models.Keys().Count() == 0 {
		privateKey, publicKey, err := utils.GenerateRSAKeys()
		if err != nil {
			panic(err)
		}

		models.Keys().Create(contracts.Fields{
			"creator_id":  user.Id,
			"name":        "default",
			"public_key":  publicKey,
			"private_key": privateKey,
		})
		logs.Default().Info(fmt.Sprintf("已创建默认密钥"))
	} else {
		logs.Default().Info(fmt.Sprintf("已存在默认密钥"))
	}

	return nil
}
