package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for our program
type Config struct {
	Port  int
	Fiber FiberConfig
	// Add other configuration fields as needed
}

func LoadConfig() (*Config, error) {
	port := 8080 // Default port

	if portStr := os.Getenv("APP_PORT"); portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid APP_PORT: %v", err)
		}
	}

	fiberConfig := DefaultFiberConfig()

	bodyLimit := os.Getenv("FIBER_BODY_LIMIT")
	if bodyLimit != "" {
		limit, err := strconv.Atoi(bodyLimit)
		if err != nil {
			return nil, fmt.Errorf("invalid FIBER_BODY_LIMIT: %v", err)
		}
		fiberConfig.BodyLimit = limit
	}

	return &Config{
		Port:  port,
		Fiber: fiberConfig,
	}, nil
}
