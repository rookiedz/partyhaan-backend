package routes

import (
	"partyhann/backend/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func AuthenRouter() *fiber.App {
	router := fiber.New()
	hdlAuthen := handlers.NewAuthen()
	router.Post("/login", hdlAuthen.Login)
	router.Post("/register", hdlAuthen.Register)
	return router
}
