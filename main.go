package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"CurrentYear": time.Now().Year(),
		})
	})

	app.Listen(":3000")
}
