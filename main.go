package main

import (
	"log"

	"go-fiber-app/package/common/config"
	"go-fiber-app/package/common/db"
	"go-fiber-app/package/routes/stations"
	"go-fiber-app/package/routes/trips"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	trips.RegisterRoutes(app, h)
	stations.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
