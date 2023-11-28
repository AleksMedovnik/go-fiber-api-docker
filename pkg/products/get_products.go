package products

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get Products
// @Tags Product
// @Description Get Product List
// @ID get-product-list
// @Accept json
// @Produce json
// @Router /products/ [get]
func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}