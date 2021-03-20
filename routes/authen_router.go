package routes

import (
	"partyhann/backend/handlers"

	fiber "github.com/gofiber/fiber/v2/"
)

func InitialAuthenRouter(router fiber.Router) {
	hdlAuthen := handlers.NewAuthen()
	router.Post("/login", hdlAuthen.Login)
	router.Post("/register", hdlAuthen.Register)
}
