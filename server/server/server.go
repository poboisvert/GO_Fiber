package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InitializeApp() {

	router := fiber.New()

	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(router.Listen(":4000"))

}
