package httpd

import (
	"partyhann/backend/routes"

	fiber "github.com/gofiber/fiber/v2/"
)

func Start() {
	app := fiber.New()
	routes.InitialRouter(app)
}
