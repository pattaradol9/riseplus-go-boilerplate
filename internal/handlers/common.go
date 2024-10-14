package handlers

import "github.com/gofiber/fiber/v2"

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Heathz(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Readyz(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
