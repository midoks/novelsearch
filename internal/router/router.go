package router

import (
	"github.com/gofiber/fiber/v2"
	//
)

func init() {

}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Home(c *fiber.Ctx) error {
	return c.Render("templates/fontends/home", fiber.Map{
		"Title": "Hello, World!",
	})
}
