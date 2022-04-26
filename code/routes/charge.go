package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/arvan/code/controllers"
	"github.com/mhd7966/arvan/code/log"
)

func UserRouter(app fiber.Router) {

	api := app.Group("/charge")

	api.Get("/:charge_code", controllers.GetChargeCode)
	api.Post("", controllers.NewChargeCode)
	api.Post("/:charge_code/apply", controllers.ApplyCharge)
	api.Post("/:charge_code/rollback", controllers.RollbackCharge)

	log.Log.Info("Charge routes created :)")

}
