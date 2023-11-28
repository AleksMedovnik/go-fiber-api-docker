package products

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get Product
// @Tags Product
// @Description Get Product
// @ID get-product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Router /products/{id} [get]
func (h handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}