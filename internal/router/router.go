package router

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Home(c *fiber.Ctx) error {
	return c.Render("templates/fontends/home", fiber.Map{
		"Title": "Hello, World!",
	})
}

func Soso(c *fiber.Ctx) error {
	return c.Render("templates/fontends/index/soso", fiber.Map{
		"Title": "Hello, World!",
	})
}
