package api

import (
	"abc/app/api/logs"
	"abc/app/api/user"
	"github.com/gofiber/fiber/v2"
)

func InitRouters(app *fiber.App) {
	logs.LogsHandlers(app)
	user.UserHandlers(app)
}
