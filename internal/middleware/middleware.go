package middleware

import (
	"riseplus-go-boilerplate/internal/config"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupMiddlewares(app *fiber.App, cfg *config.Config) {
	// Add all middlewares in a single call to Use
	app.Use(
		fiberzap.New(cfg.FiberZapConfig()),
		cors.New(),
		helmet.New(),
		recover.New(),
		requestid.New(),
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
	)
}
