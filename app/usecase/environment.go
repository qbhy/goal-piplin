package usecase

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
)

func CreateEnvironment(name string, projectId int) (*models.ProjectEnvironment, error) {
	return models.ProjectEnvironments().CreateE(contracts.Fields{
		"project_id": projectId,
		"name":       name,
		"settings": models.EnvironmentSettings{
			Servers:  make([]models.Server, 0),
			Cabinets: make([]string, 0),
		},
	})
}
