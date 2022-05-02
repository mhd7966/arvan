package controllers

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/wallet/constants"
	"github.com/mhd7966/arvan/wallet/log"
	"github.com/mhd7966/arvan/wallet/models"
	repositories "github.com/mhd7966/arvan/wallet/repositoies"
	"github.com/sirupsen/logrus"
)

// Get godoc
// @Summary get user
// @Description return user
// @ID get_user
// @Tags User
// @Param phone_number path string true "phone number of user"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /users/{phone_number} [get]
func GetUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"
	phoneNumber := c.Params("phone_number")

	if valid := ValidationPhoneNumber(phoneNumber); !valid {
		response.Message = "Not Valid PhoneNumber "
		log.Log.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
		}).Info("GetUser. PhoneNumber isn't valid !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUser(phoneNumber); !existUser {
		response.Message = "User Not Found"
		log.Log.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
			"exist_code":   existUser,
		}).Info("GetUser. User doesn't exist !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := repositories.GetUser(phoneNumber)
	if err != nil {
		response.Message = "Get User Failed!"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("GetUser. Get User Have Error!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	response.Data = user
	log.Log.WithFields(logrus.Fields{
		"user":     user,
		"response": response.Message,
	}).Info("GetUser. Get َUser Successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

func ValidationPhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(constants.PHONENUMBER)
	return re.MatchString(phoneNumber)
}
