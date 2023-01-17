package app

import (
	"abc/app/api"
	"abc/app/storage"
	"github.com/gofiber/fiber/v2"
	"log"
)

func App() {
	app := fiber.New()

	storage.InitDB()
	defer func() {
		err := storage.CloseDB()
		if err != nil {
			log.Fatal("Error of database closing")
		}
	}()
	storage.Migrate()

	api.InitRouters(app)

	log.Fatal(app.Listen(":3000"))
}
