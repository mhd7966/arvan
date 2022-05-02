package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req"
	"github.com/mhd7966/arvan/wallet/configs"
	"github.com/mhd7966/arvan/wallet/constants"
	"github.com/mhd7966/arvan/wallet/inputs"
	"github.com/mhd7966/arvan/wallet/log"
	"github.com/mhd7966/arvan/wallet/models"
	repositories "github.com/mhd7966/arvan/wallet/repositoies"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

// charge code godoc
// @Summary charge
// @Description charge code
// @ID charge
// @Tags Transaction
// @Param charge body inputs.Charge true "charge body"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /transactions [post]
func Charge(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"

	charge := new(inputs.Charge)
	if err := c.BodyParser(charge); err != nil {
		response.Message = "Parse Body Failed"
		log.Log.WithFields(logrus.Fields{
			"response": response.Message,
			"error":    err.Error(),
		}).Error("Charge. Parse body to charge failed!")
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	if valid := ValidationPhoneNumber(charge.PhoneNumber); !valid {
		response.Message = "Not Valid PhoneNumber "
		log.Log.WithFields(logrus.Fields{
			"phone_number": charge.PhoneNumber,
		}).Info("Charge. PhoneNumber isn't valid !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUser(charge.PhoneNumber); !existUser {
		_, err := repositories.CreateUser(charge.PhoneNumber)
		if err != nil {
			response.Message = "Create User Failed"
			log.Log.WithFields(logrus.Fields{
				"charge_body": charge,
				"response":    response.Message,
				"error":       err.Error(),
			}).Error("Charge. Create user Failed!")
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
	}

	existTransaction, err := repositories.ExistTransaction(charge.PhoneNumber, charge.Code)
	if err != nil {
		response.Message = "Check Transaction Info Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_body": charge,
			"response":    response.Message,
			"error":       err.Error(),
		}).Error("Charge. There is a error in check exist query!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existTransaction {
		response.Message = "Duplicate Charge"
		log.Log.WithFields(logrus.Fields{
			"phone_number": charge.PhoneNumber,
			"exist_code":   existTransaction,
		}).Info("Charge. User try duplicate code !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	chargeCode, err := GetChargeCodeReq(charge.Code)
	if err != nil {
		RollbackCapacityReq(charge.Code)
		response.Message = err.Error()
		log.Log.WithFields(logrus.Fields{
			"charge_code": charge.Code,
			"error":       err,
		}).Error("Charge. Code Charge Request Failed !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := repositories.GetUser(charge.PhoneNumber)
	if err != nil {
		RollbackCapacityReq(charge.Code)
		response.Message = "Get User Failed"
		log.Log.WithFields(logrus.Fields{
			"charge_code": charge.Code,
			"error":       err,
		}).Error("Charge. Get User Failed !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	transaction := models.Transaction{
		UserID:      int(user.ID),
		Code:        chargeCode.Name,
		CodeType:    constants.CHARGE,
		Value:       chargeCode.Value,
		ValueType:   constants.INCREASE,
		InitBalance: user.Balance,
		NewBalance:  user.Balance + chargeCode.Value,
	}

	if err = repositories.CreateChargeTransaction(&transaction, user); err != nil {
		RollbackCapacityReq(charge.Code)
		response.Message = "Create Transaction Failed"
		log.Log.WithFields(logrus.Fields{
			"transaction": transaction,
			"user":        user,
			"error":       err,
		}).Error("Charge. Create Transaction and Update Balance Failed !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = transaction
	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"transaction": transaction,
		"user":        user,
		"response":    response.Message,
	}).Info("Charge. Charge Code successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

// get transactions godoc
// @Summary get transactions
// @Description get transactions
// @ID get_transactions
// @Tags Transaction
// @Param phone_number path string true "phone number of user"
// @Success 200 {object} models.Response
// @Failure 400 json httputil.HTTPError
// @Router /transactions/{phone_number} [get]
func GetTransactions(c *fiber.Ctx) error {

	var response models.Response
	response.Status = "error"
	phoneNumber := c.Params("phone_number")

	if valid := ValidationPhoneNumber(phoneNumber); !valid {
		response.Message = "Not Valid PhoneNumber "
		log.Log.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
		}).Info("Charge. PhoneNumber isn't valid !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if existUser := repositories.ExistUser(phoneNumber); !existUser {
		response.Message = "User Not Found"
		log.Log.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
			"exist_code":   existUser,
		}).Info("getTransactions. User doesn't exist !")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	transactions, err := repositories.GetTransactions(phoneNumber)
	if err != nil {
		response.Message = "Get Transaction History Failed"
		log.Log.WithFields(logrus.Fields{
			"phone_number": phoneNumber,
			"response":     response.Message,
			"error":        err.Error(),
		}).Error("getTransactions. There is no contact for this request!")
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Data = transactions
	response.Message = "OK!"
	response.Status = "succes"
	log.Log.WithFields(logrus.Fields{
		"transactions": transactions,
		"response":     response.Message,
	}).Info("getTransactions. Get transaction history successful :)")
	return c.Status(fiber.StatusOK).JSON(response)

}

func GetChargeCodeReq(chargeCode string) (*models.ChargeCode, error) {

	config := configs.Cfg.CodeService

	header := req.Header{
		"Host":   config.Host,
		"Origin": "http://" + config.Host,
	}

	url := "http://" + config.Host + "/v0/charge/" + chargeCode + "/apply"
	r, err := req.Post(url, header)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"URL":    url,
			"Header": header,
			"error":  err.Error(),
		}).Debug("GetChargeCodeReq. Get charge code request failed!")
		return nil, err
	}

	log.Log.Debug("GetChargeCodeReq. Send request ok :)")
	var resp models.Response
	err = r.ToJSON(&resp)

	if err != nil {
		log.Log.Debug("GetChargeCodeReq. convert response model failed!")
		return nil, err
	}

	var chargeCodeObj models.ChargeCode
	if resp.Status == "error" {
		log.Log.WithFields(logrus.Fields{
			"Response": r.String(),
		}).Debug("GetChargeCodeReq. Response of charge code appply request is incorrect!")
		return nil, errors.New(resp.Message)

	} else {
		chargeCodeMap := resp.Data.(map[string]interface{})
		mapstructure.Decode(chargeCodeMap, &chargeCodeObj)
		log.Log.WithFields(logrus.Fields{
			"Response":   r.String(),
			"chargeCode": chargeCodeObj,
		}).Debug("GetChargeCodeReq. Response of charge code appply request is correct!")
		return &chargeCodeObj, nil
	}

}

func RollbackCapacityReq(chargeCode string) error {

	config := configs.Cfg.CodeService

	header := req.Header{
		"Host":   config.Host,
		"Origin": "http://" + config.Host,
	}

	url := "http://" + config.Host + "/v0/charge/" + chargeCode + "/rollback"
	r, err := req.Post(url, header)
	if err != nil {
		log.Log.WithFields(logrus.Fields{
			"URL":    url,
			"Header": header,
			"error":  err.Error(),
		}).Debug("GetChargeCodeReq. rollback capacity charge code request failed!")
		return err
	}

	log.Log.Debug("GetChargeCodeReq. Send request ok :)")
	var resp models.Response
	err = r.ToJSON(&resp)

	if err != nil {
		log.Log.Debug("GetChargeCodeReq. convert response model failed!")
		return err
	}

	if resp.Status == "error" {
		log.Log.WithFields(logrus.Fields{
			"Response": r.String(),
		}).Debug("GetChargeCodeReq. Response of rollback charge code request is incorrect!")
		return errors.New(resp.Message)

	} else {
		log.Log.WithFields(logrus.Fields{
			"Response": r.String(),
		}).Debug("GetChargeCodeReq. Response of rollback charge code  request is correct!")
		return nil
	}

}
