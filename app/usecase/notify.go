package usecase

import (
	"fmt"
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal-piplin/app/models"
	"sync"
)

var (
	sse     contracts.SseFactory
	sseOnce sync.Once
)

func getSse() contracts.SseFactory {
	sseOnce.Do(func() {
		sse = application.Get("sse.factory").(contracts.SseFactory)
	})
	return sse
}

func DeploymentNotify(deployment *models.Deployment) {
	err := getSse().Sse("/api/notify").Broadcast(deployment)
	if err != nil {
		fmt.Println(err.Error())
	}
}
