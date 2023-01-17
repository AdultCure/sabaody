package user

import "github.com/gofiber/fiber/v2"

func UserHandlers(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/get-all", GetAllUsers)

}

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString(c.Params("test"))
}
