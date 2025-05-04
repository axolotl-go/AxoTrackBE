package routes

import "github.com/gofiber/fiber/v2"

func HelloApp(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}