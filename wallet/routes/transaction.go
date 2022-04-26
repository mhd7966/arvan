package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/wallet/controllers"
	"github.com/mhd7966/arvan/wallet/log"
)

func TransactionRouter(app fiber.Router) {

	api := app.Group("/transaction")

	api.Post("/", controllers.Charge)
	api.Get("/:phone_number", controllers.History)

	log.Log.Info("Transaction routes created :)")

}
