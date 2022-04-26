package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/code/log"
)

func MainRouter(app fiber.Router) {
	api := app.Group("/v0")

	UserRouter(api)

	log.Log.Info("All routes created successfully :)")

}
