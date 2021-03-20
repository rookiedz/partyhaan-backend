package routes

import fiber "github.com/gofiber/fiber/v2/"

func InitialRouter(app *fiber.App) {
	api := app.Group("/api")
	InitialAuthenRouter(api)
	InitialJoinRouter(api)
	InitialPartyRouter(api)
	InitialUserRouter(api)
}
