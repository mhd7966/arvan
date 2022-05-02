package repositories

import (
	"time"

	"github.com/mhd7966/arvan/code/connections"
	"github.com/mhd7966/arvan/code/inputs"
	"github.com/mhd7966/arvan/code/log"
	"github.com/mhd7966/arvan/code/models"
	"github.com/sirupsen/logrus"
)

func GetCharge(chargeCode string) (*models.ChargeCode, error) {
	var charge models.ChargeCode
	if result := connections.DB.Debug().Where("name = ?", chargeCode).Find(&charge); result.Error != nil {
		return nil, result.Error
	}

	return &charge, nil
}

func UpdateCapacity(charge *models.ChargeCode) error {

	if result := connections.DB.Save(charge); result.Error != nil {
		return result.Error
	}

	return nil
}

func ExistChargeCode(chargeCode string) bool {
	var charge models.ChargeCode

	if result := connections.DB.Where(models.ChargeCode{Name: chargeCode}).First(&charge); result.RowsAffected == 0 {
		return false
	}
	return true
}

func CreateChargeCode(chargeCodeBody inputs.ChargeCode) (*models.ChargeCode, error) {
	exTime, err := time.Parse(time.RFC3339, chargeCodeBody.ExpirationDate)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"charge": chargeCodeBody.ExpirationDate,
			"format": time.RFC3339,
			"error":  err.Error(),
		}).Error("Repo.CreateChargeCode. Convert string to time filed!")
		return nil, err
	}
	chargeCode := models.ChargeCode{
		Name:           chargeCodeBody.Name,
		Value:          chargeCodeBody.Value,
		MaxCapacity:    chargeCodeBody.MaxCapacity,
		ExpirationDate: exTime,
	}

	if result := connections.DB.Create(&chargeCode); result.Error != nil {
		return nil, result.Error
	}

	return &chargeCode, nil
}
