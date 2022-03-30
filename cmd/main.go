package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/go-fiber-api/pkg/common/config"
	"github.com/hellokvn/go-fiber-api/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db.Init(c.DBUrl)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(c.Port)
}
