package main

import (
	"app/db"
	"app/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	server := fiber.New()
	handler.RouteProducts(server)
	handler.RouteAuth(s

	server.Listen(":8080")
}
