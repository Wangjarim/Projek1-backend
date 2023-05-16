package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	belajarRoutes "sistem-informasi-klinik/internal/routes/belajar"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	// Setup the Node Routes
	belajarRoutes.SetupBelajarRoutes(api)
}
