package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"

	_ "go-fiber-api-docker/docs"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	routes := app.Group("/products")
	routes.Post("/", h.AddProduct)
	routes.Get("/", h.GetProducts)
	routes.Get("/:id", h.GetProduct)
	routes.Delete("/:id", h.DeleteProduct)
	routes.Put("/:id", h.UpdateProduct)
}