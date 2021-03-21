package httpd

import (
	"partyhaan/backend/config"
	"partyhaan/backend/routes"
	"partyhaan/backend/stores/mariadb"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start() {
	c := config.LoadConfiguration("config.json")
	mariadb.QueryConnect(c.MariaDB.DSNQuery)
	mariadb.ExecConnect(c.MariaDB.DSNExec)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")
	api.Mount("/users", routes.UserRouter())
	api.Mount("/parties", routes.PartyRouter())

	app.Listen(":8080")
}
