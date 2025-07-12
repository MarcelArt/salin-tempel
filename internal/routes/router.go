package routes

import (
	"github.com/MarcelArt/salin-tempel/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	clipboard := app.Group("/clipboard")
	clipboard.Post("/", handlers.Copy)
	clipboard.Post("/img", handlers.CopyPicture)

	qr := app.Group("/qr")
	qr.Get("/", handlers.QRCode)
}
