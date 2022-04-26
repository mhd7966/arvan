package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/wallet/controllers"
	"github.com/mhd7966/arvan/wallet/log"
)

func UserRouter(app fiber.Router) {

	api := app.Group("/user")

	api.Get("/:phone_number", controllers.GetUser)

	log.Log.Info("User routes created :)")

}
