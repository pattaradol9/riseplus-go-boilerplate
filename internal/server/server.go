package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"riseplus-go-boilerplate/internal/config"
	"riseplus-go-boilerplate/internal/routes"
	"sync"
	"syscall"
	"time"

	"riseplus-go-boilerplate/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	config *config.Config
}

func New(cfg *config.Config) *Server {
	app := fiber.New(cfg.FiberConfig())

	// Apply middlewares
	middleware.SetupMiddlewares(app, cfg)

	routes.SetupRoutes(app)

	return &Server{
		app:    app,
		config: cfg,
	}
}

func (s *Server) Start() error {
	// Use a WaitGroup to ensure all goroutines finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Start server
	go func() {
		defer wg.Done()
		if err := s.app.Listen(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != fiber.ErrServiceUnavailable {
			s.config.Fiber.Logger.Error("Server error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.config.Fiber.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}

	// Wait for the server goroutine to finish
	wg.Wait()

	s.config.Fiber.Logger.Info("Server exited gracefully")
	return nil
}
