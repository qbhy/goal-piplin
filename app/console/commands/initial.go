package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/supports/commands"
	"github.com/goal-web/supports/logs"
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
	if user := models.Users().Where("username", username).First(); user != nil {
		logs.Default().Info("piplin 用户已存在")
		models.Users().Where("id", user.Id).Update(contracts.Fields{
			"password": cmd.hash.Make(password, nil),
		})
		logs.Default().Info(fmt.Sprintf("已将密码重置为 %s", password))
	}

	models.Users().Create(contracts.Fields{
		"username": username,
		"nickname": username,
		"avatar":   "",
		"role":     "system",
		"password": cmd.hash.Make(password, nil),
	})
	logs.Default().Info(fmt.Sprintf("已创建用户 %s 密码为 %s", username, password))

	return nil
}
