package main

import (
	"github.com/mhd7966/arvan/code/configs"
	"github.com/mhd7966/arvan/code/connections"
	_ "github.com/mhd7966/arvan/code/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/code/log"
	"github.com/mhd7966/arvan/code/routes"
)

// @title Code API
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
		Prefork: true, //number of core cpu -> create version
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	log.Log.Info("Swagger handler route created :)")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi Stupid:)")
	})

	routes.MainRouter(app)

	app.Listen(":3001")

}
