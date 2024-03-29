package console

import (
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/console/commands"
)

func NewService() contracts.ServiceProvider {
	return console.NewService(NewKernel)
}

func NewKernel(app contracts.Application) contracts.Console {
	return &Kernel{console.NewKernel(app, []contracts.CommandProvider{
		commands.Runner,
		commands.NewHello,
		commands.NewInitial,
	}), app}
}

type Kernel struct {
	*console.Kernel
	app contracts.Application
}

func (kernel *Kernel) Schedule(schedule contracts.Schedule) {
	//schedule.Call(func() {
	//	logs.Default().Info("周日每5秒钟打印 周日愉快")
	//}).EveryFiveSeconds().Sundays()
}
