package logs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func LogsHandlers(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Sever Metrics Page"}))
}
