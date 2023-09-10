package commands

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"github.com/goal-web/supports/logs"
)

func NewInitial(app contracts.Application) contracts.Command {
	return &Initial{
		Command: commands.Base("init", "打印 hello goal"),
	}
}

type Initial struct {
	commands.Command
}

func (hello Initial) Handle() any {
	logs.Default().Info("hello goal " + hello.GetString("say"))
	return nil
}
