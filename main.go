package main

import (
	"abc/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log"
)

func main() {
	connectResult, err := storage.Storage()
	if err != nil {
		fmt.Printf(connectResult)
	}

	app := fiber.New()

	admin := app.Group("/preference")
	admin.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(ctx.Params("test"))
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Sever Metrics Page"}))

	log.Fatal(app.Listen(":3000"))
}
