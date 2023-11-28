package main

import (
	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/products"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Go-fiber-api-docker
// @version 1.0
// @description This is an example of a REST API 
// @host localhost:3000
// @BasePath /
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()
	app.Use(cors.New())

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}