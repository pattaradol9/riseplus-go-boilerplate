package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	app := fiber.New()
	app.Get("/", HelloWorld)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, World!", string(body))
}

func TestHeathz(t *testing.T) {
	app := fiber.New()
	app.Get("/healthz", Heathz)

	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestReadyz(t *testing.T) {
	app := fiber.New()
	app.Get("/readyz", Readyz)

	req := httptest.NewRequest("GET", "/readyz", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
