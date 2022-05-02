package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/code/inputs"
	"github.com/mhd7966/arvan/code/log"
	"github.com/mhd7966/arvan/code/models"
	"github.com/mhd7966/arvan/code/repositoies"
	"github.com/sirupsen/logrus"
)

// get charge godoc
// @Summary get charge
// @Description get charge
// @ID get_charge
// @Tags Charge
// @Param charge_code path string true "charge code name"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /charge/{charge_code} [get]
func GetChargeCode(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"
	chargeCode := c.Params("charge_code")

	if exist := repositories.ExistChargeCode(chargeCode); !exist {
		response.Message = "Charge Code Not Found"
		log.Log.WithFields(logrus.Fields{
			"charge_code": chargeCode,
			"response":    response.Message,
		}).Error("GetChargeCode. The charge code doesn't exist!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	charge, err := repositories.GetCharge(chargeCode)
	if err != nil {
		response.Message = "Get Charge Code Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_code": chargeCode,
			"response":    response.Message,
			"error":       err.Error(),
		}).Error("GetCharge. There is a error in get charge query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = charge
	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"logs":     charge,
		"response": response.Message,
	}).Info("GetCharge.Get charge successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// create charge godoc
// @Summary create charge
// @Description create charge
// @ID create_charge
// @Tags Charge
// @Param charge_code body inputs.ChargeCode true "charge code info"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /charge [post]
func NewChargeCode(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"

	chargeCodeBody := new(inputs.ChargeCode)
	if err := c.BodyParser(chargeCodeBody); err != nil {
		response.Message = "Parse Body Failed"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("NewChargeCode. Parse body to ChargeCode body failed!")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	validate := validator.New()
	if err := validate.Struct(chargeCodeBody); err != nil {
		validateError := "error : "
		for _, err := range err.(validator.ValidationErrors) {
			validateError += err.StructField() + ", "
		}
		response.Message = "Validate Charge Info Failed"
		response.Data = validateError
		log.Log.WithFields(logrus.Fields{
			"response":       response.Message,
			"validate_error": validateError,
			"error":          err.Error(),
		}).Error("NewUser. Validate charge info failed!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if exist := repositories.ExistChargeCode(chargeCodeBody.Name); exist {
		response.Message = "Duplicate Charge Code"
		log.Log.WithFields(logrus.Fields{
			"charge_code_body": chargeCodeBody,
			"response":         response.Message,
		}).Error("NewChargeCode. The charge code is duplicate!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	chargeCode, err := repositories.CreateChargeCode(*chargeCodeBody)
	if err != nil {
		response.Message = "Create Charge Code Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_code_body": chargeCodeBody,
			"response":         response.Message,
			"error":            err.Error(),
		}).Error("NewChargeCode. There is a error in create charge query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = chargeCode
	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"charge_code_body": chargeCodeBody,
		"charge_code":      chargeCode,
		"response":         response.Message,
	}).Info("NewChargeCode.Create charge successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// charge code godoc
// @Summary charge
// @Description charge code
// @ID charge
// @Tags Charge
// @Param charge_code path string true "charge code name"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /charge/{charge_code}/apply [post]
func ApplyCharge(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"
	chargeCode := c.Params("charge_code")

	charge, err := repositories.GetCharge(chargeCode)
	if err != nil {
		response.Message = "Get Charge Code Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_code": chargeCode,
			"response":    response.Message,
			"error":       err.Error(),
		}).Error("ApplyCharge. There is a error in get charge query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if charge.MaxCapacity > charge.Capacity && time.Now().Before(charge.ExpirationDate) {
		charge.Capacity++
		if err := repositories.UpdateCapacity(charge); err != nil {
			response.Message = "Update Charge Capacity Failed"
			log.Log.WithFields(logrus.Fields{
				"charge":   charge,
				"response": response.Message,
				"error":    err.Error(),
			}).Error("ApplyCharge. There is a error in update charge capacity query!")
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
	} else {
		response.Message = "The code is not usable"
		return c.Status(fiber.StatusNotAcceptable).JSON(response)

	}

	response.Data = *charge
	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"charge_data": response.Data,
		"response":    response.Message,
	}).Info("ApplyCharge. Charge successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// charge rollback godoc
// @Summary charge
// @Description charge rollback
// @ID charge_rollback
// @Tags Charge
// @Param charge_code path string true "charge code name"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /charge/{charge_code}/rollback [post]
func RollbackCharge(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"
	chargeCode := c.Params("charge_code")

	if exist := repositories.ExistChargeCode(chargeCode); !exist {
		response.Message = " Charge Code Not Exist"
		log.Log.WithFields(logrus.Fields{
			"charge_code": chargeCode,
			"response":    response.Message,
		}).Error("RollbackCharge. charge code not exist !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	charge, err := repositories.GetCharge(chargeCode)
	if err != nil {
		response.Message = "Get Charge Code Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_code": chargeCode,
			"response":    response.Message,
			"error":       err.Error(),
		}).Error("RollbackCharge. There is a error in get charge query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	charge.Capacity--
	if err = repositories.UpdateCapacity(charge); err != nil {
		response.Message = "Update Charge Capacity Failed"
		log.Log.WithFields(logrus.Fields{
			"charge":   charge,
			"response": response.Message,
			"error":    err.Error(),
		}).Error("RollbackCharge. There is a error in update charge capacity query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"logs":     charge,
		"response": response.Message,
	}).Info("RollbackCharge. Rollback charge successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}
