package usecase

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
	"github.com/goal-web/supports/utils"
	"github.com/savsgio/gotils/uuid"
)

func CreateProject(fields contracts.Fields) (models.Project, error) {
	var project models.Project
	if utils.ToInt(fields["key_id"], 0) == 0 {
		key, err := CreateKey(utils.ToString(fields["name"], ""))
		fields["key_id"] = key.Id
		if err != nil {
			return project, err
		}
	}

	fields["uuid"] = uuid.V4()
	fields["settings"] = models.ProjectSettings{}

	project = models.Projects().Create(fields)
	return project, nil
}

func GetProjectDetail(id any) models.ProjectDetail {
	project := models.Projects().Find(id)
	return models.ProjectDetail{
		Project: project,
		Key:     models.Keys().Find(project.KeyId),
		Group:   models.Groups().Find(project.GroupId),
	}
}
