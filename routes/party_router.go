package routes

import (
	"partyhaan/backend/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func PartyRouter() *fiber.App {
	router := fiber.New()
	hdlParty := handlers.NewParty()
	router.Get("/:id/users", hdlParty.Users)
	router.Get("/:id", hdlParty.Find)
	router.Put("/:id", hdlParty.Update)
	router.Delete("/:id", hdlParty.Delete)
	router.Get("/", hdlParty.All)
	router.Post("/", hdlParty.Crete)
	return router
}
