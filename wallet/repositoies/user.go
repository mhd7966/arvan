package repositories

import (
	"github.com/mhd7966/arvan/wallet/connections"
	"github.com/mhd7966/arvan/wallet/models"
)

func ExistUser(phoneNumber string) bool {
	var user models.User

	if result := connections.DB.Where(models.User{PhoneNumber: phoneNumber}).First(&user); result.RowsAffected == 0 {
		return false
	}
	return true
}

func GetUser(phoneNumber string) (*models.User, error) {
	var user models.User

	if result := connections.DB.Where(models.User{PhoneNumber: phoneNumber}).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(phoneNumber string) (*models.User, error) {
	user := models.User{
		PhoneNumber: phoneNumber,
	}

	if result := connections.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func UpdateBalance(user models.User) error {

	if result := connections.DB.Model(&user).Updates(models.User{Balance: user.Balance}); result.Error != nil {
		return result.Error
	}

	return nil
}
