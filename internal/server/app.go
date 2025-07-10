package server

import (
	"github.com/MarcelArt/salin-tempel/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Run() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", helloWorld)

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
