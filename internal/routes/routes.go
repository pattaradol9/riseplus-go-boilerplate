package routes

import (
	"riseplus-go-boilerplate/internal/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "RisePlus Go Boilerplate",
		Refresh: 2 * time.Second,
	}))
	app.Get("/", handlers.HelloWorld)
	app.Get("/healthz", handlers.Heathz)
	app.Get("/readyz", handlers.Readyz)

	// Add more routes here as needed
	// For example:
	// app.Get("/api/users", handlers.GetUsers)
	// app.Post("/api/users", handlers.CreateUser)
}
