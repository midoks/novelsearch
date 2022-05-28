package admin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {

	fmt.Println("dd")
	return c.Render("templates/backends/index/index", fiber.Map{
		"Title":     "Hello, World!",
		"adminPath": "admin",
	})
}
