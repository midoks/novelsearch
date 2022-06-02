package admin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/midoks/novelsearch/internal/mgdb"
)

func Admin(c *fiber.Ctx) error {
	fmt.Println("dd")
	return c.Render("templates/backends/index/index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func Login(c *fiber.Ctx) error {
	return c.Render("templates/backends/login/index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func SpiderList(c *fiber.Ctx) error {

	result, err := mgdb.NovelSourceSearch("", "-")
	if err != nil {
		return err
	}

	return c.Render("templates/backends/spider/index", fiber.Map{
		"menu_title":     "爬虫设置",
		"menu_sub_title": "列表",
		"list":           result,
	})
}

func SpiderAdd(c *fiber.Ctx) error {

	m := fiber.Map{
		"menu_title":     "爬虫设置",
		"menu_sub_title": "添加",
	}

	return c.Render("templates/backends/spider/add", m)
}
