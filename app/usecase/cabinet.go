package usecase

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/example/app/models"
	"github.com/golang-module/carbon/v2"
)

func CreateCabinet(name string, settings any) (*models.Cabinet, error) {
	if models.Cabinets().Where("name", name).Count() > 0 {
		return nil, errors.New("机柜已存在")
	}

	cabinet, err := models.Cabinets().CreateE(contracts.Fields{
		"name":       name,
		"settings":   settings,
		"created_at": carbon.Now().ToDateTimeString(),
		"updated_at": carbon.Now().ToDateTimeString(),
	})
	if err != nil {
		return nil, err
	}

	if err = verifyCabinet(cabinet); err != nil {
		models.Cabinets().Where("id", cabinet.Id).Delete()
		return nil, err
	}

	return cabinet, nil
}

func verifyCabinet(cabinet *models.Cabinet) contracts.Exception {

	return nil
}

func UpdateCabinet(id any, name string, settings any) error {
	if models.Cabinets().Where("id", "!=", id).Where("name", name).Count() > 0 {
		return errors.New("机柜名称已存在")
	}

	cabinet := models.CabinetClass.New(contracts.Fields{
		"settings": settings,
	})

	if err := verifyCabinet(&cabinet); err != nil {
		return err
	}

	_, err := models.Cabinets().Where("id", id).UpdateE(contracts.Fields{
		"name":       name,
		"settings":   settings,
		"updated_at": carbon.Now().ToDateTimeString(),
	})

	return err
}

func DeleteCabinet(id any) error {
	_, err := models.Cabinets().WhereIn("id", id).DeleteE()
	return err
}
