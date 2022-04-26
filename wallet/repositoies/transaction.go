package repositories

import (
	"github.com/mhd7966/arvan/wallet/connections"
	"github.com/mhd7966/arvan/wallet/log"
	"github.com/mhd7966/arvan/wallet/models"
)

func ExistTransaction(phoneNumber string, code string) (bool, error) {
	var transaction models.Transaction

	if result := connections.DB.Joins("JOIN users on transactions.user_id=users.id").
		Where("users.phone_number=?", phoneNumber).Where("transactions.code=?", code).Find(&transaction); result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func GetTransactions(phoneNumber string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if result := connections.DB.Joins("JOIN users on transactions.user_id=users.id").
		Where("users.phone_number=?", phoneNumber).Find(&transactions); result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func CreateChargeTransaction(transaction *models.Transaction, user *models.User) error {

	tx := connections.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Log.WithField("err", r).Debug("CreateChargeTransaction. Recover Failed!")
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Log.WithField("err", err).Debug("CreateChargeTransaction. Transaction Have Error!")
		tx.Rollback()
		return err
	}

	if result := tx.Create(&transaction); result.Error != nil {
		log.Log.WithField("err", result.Error).Debug("CreateChargeTransaction. Create Transaction Have Error!")
		tx.Rollback()
		return result.Error
	}

	user.Balance += transaction.Value
	if result := tx.Save(user); result.Error != nil {
		user.Balance -= transaction.Value
		log.Log.WithField("err", result.Error).Debug("CreateChargeTransaction. Update Balance Have Error!")
		tx.Rollback()
		return result.Error
	}

	err := tx.Commit().Error
	if err != nil {
		log.Log.WithField("err", err).Debug("CreateChargeTransaction. Commit Transaction Have Error!")
		return err
	}

	return nil
}
