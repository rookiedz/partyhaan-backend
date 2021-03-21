package routes

import fiber "github.com/gofiber/fiber/v2"

func InitialJoinRouter() *fiber.App {
	router := fiber.New()
	return router
}
