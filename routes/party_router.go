package routes

import (
	"partyhann/backend/handlers"

	fiber "github.com/gofiber/fiber/v2/"
)

func InitialPartyRouter(router fiber.Router) {
	hdlParty := handlers.NewParty()
	router.Get("/parties/:id/users", hdlParty.Users)
	router.Get("/parties/:id", hdlParty.Find)
	router.Put("/parties/:id", hdlParty.Update)
	router.Delete("/parties/:id", hdlParty.Delete)
	router.Get("/parties", hdlParty.All)
	router.Post("/parties", hdlParty.Crete)
}
