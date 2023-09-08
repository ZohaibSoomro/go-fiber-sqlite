package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/zohaibsoomro/go-fiber-sqlite/db"
	"github.com/zohaibsoomro/go-fiber-sqlite/routes"
)

func main() {
	app := fiber.New()
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(80))
}
