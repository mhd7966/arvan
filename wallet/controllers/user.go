package controllers

import (
	"github.com/gofiber/fiber/v2"
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
// @Router /user/{phone_number} [get]
func GetUser(c *fiber.Ctx) error {
	var response models.Response
	response.Status = "error"
	phoneNumber := c.Params("phone_number")

	user, err := repositories.GetUser(phoneNumber)
	if err != nil {
		response.Message = "This User Have Error!"
		log.Log.WithFields(logrus.Fields{
			"response":   response.Message,
			"error":      err.Error(),
		}).Error("GetUser. Get User Have Error!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	response.Data = user
	log.Log.WithFields(logrus.Fields{
		"user": user,
		"response":    response.Message,
	}).Info("GetUser. Get َUser Successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}


