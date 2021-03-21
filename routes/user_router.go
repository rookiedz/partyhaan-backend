package routes

import (
	"partyhaan/backend/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func UserRouter() *fiber.App {
	router := fiber.New()
	hdlUser := handlers.NewUser()
	router.Post("/authenticate", hdlUser.Authenticate)
	router.Post("/register", hdlUser.Register)
	return router
}
