package main

import (
	"github.com/mhd7966/arvan/wallet/configs"
	"github.com/mhd7966/arvan/wallet/connections"
	_ "github.com/mhd7966/arvan/wallet/docs"
	"github.com/mhd7966/arvan/wallet/log"
	"github.com/mhd7966/arvan/wallet/routes"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

//// @host localhost:3000 -> if set when you have domain you should set domain and then you want to test it localy -> change the host -> nazarim behtare

// @title Wallet API
// @version 1.0
// @description I have no specific description
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @BasePath /v0

func main() {

	configs.SetConfig()
	log.LogInit()
	connections.ConnectDatabase()
	defer connections.CloseDB()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	log.Log.Info("Swagger handler route created :)")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi Stupid:)")
	})

	routes.MainRouter(app)

	app.Listen(":3000")

}
