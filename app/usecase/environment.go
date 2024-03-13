package usecase

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/exceptions"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal-piplin/app/models"
	"net"
	"time"
)

func CreateEnvironment(name string, projectId int) (*models.ProjectEnvironment, error) {
	if models.ProjectEnvironments().
		Where("project_id", projectId).
		Where("name", name).
		Count() > 0 {
		return nil, errors.New("环境名称已存在！")
	}

	return models.ProjectEnvironments().CreateE(contracts.Fields{
		"project_id": projectId,
		"name":       name,
		"settings": models.EnvironmentSettings{
			Servers:  make([]models.Server, 0),
			Cabinets: make([]string, 0),
		},
	})
}

func UpdateEnvironment(id any, name string, settings any) error {
	env := models.ProjectEnvironments().Find(id)
	if models.ProjectEnvironments().
		Where("project_id", env.ProjectId).
		Where("id", "!=", id).
		Where("name", name).
		Count() > 0 {
		return errors.New("环境名称已存在！")
	}

	newEnv := models.ProjectEnvironmentClass.New(contracts.Fields{
		"settings": settings,
	})

	if err := verifyEnvironment(&newEnv.Settings); err != nil {
		return err
	}

	_, err := models.ProjectEnvironments().Where("id", id).UpdateE(contracts.Fields{
		"name":       name,
		"settings":   settings,
		"updated_at": carbon.Now().ToDateTimeString(),
	})

	return err
}

func verifyEnvironment(env *models.EnvironmentSettings) contracts.Exception {
	for _, setting := range env.Servers {
		target := fmt.Sprintf("%s:%d", setting.Host, setting.Port)
		// 尝试建立连接
		conn, err := net.DialTimeout("tcp", target, 5*time.Second)
		if err != nil {
			// 如果出现错误，表示连接失败
			return exceptions.New(fmt.Sprintf("连接失败: %v\n", err))
		}
		// 关闭连接
		_ = conn.Close()
	}
	return nil
}

func DeleteEnvironment(id any) error {
	_, err := models.ProjectEnvironments().WhereIn("id", id).DeleteE()
	return err
}
