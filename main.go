package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"say": "Hi, Go fiber",
		})
	})

	app.Listen(":3000")
}
