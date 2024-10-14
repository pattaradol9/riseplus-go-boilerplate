package config

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// FiberConfig holds the configuration for the Fiber app
type FiberConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Logger       *zap.Logger
	BodyLimit    int
}

// DefaultFiberConfig returns the default Fiber configuration
func DefaultFiberConfig() FiberConfig {
	logger, _ := zap.NewProduction()
	return FiberConfig{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Logger:       logger,
		BodyLimit:    10 * 1024 * 1024, // 10 MB (default)
	}
}

// FiberConfig returns the Fiber configuration
func (c *Config) FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout:  c.Fiber.ReadTimeout,
		WriteTimeout: c.Fiber.WriteTimeout,
		IdleTimeout:  c.Fiber.IdleTimeout,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		AppName:      "RisePlus Go Boilerplate",
		BodyLimit:    c.Fiber.BodyLimit,
	}
}

// FiberZapConfig returns the Fiberzap configuration
func (c *Config) FiberZapConfig() fiberzap.Config {
	return fiberzap.Config{
		Logger: c.Fiber.Logger,
		SkipBody: func(c *fiber.Ctx) bool {
			return c.Method() == "GET"
		},
	}
}
