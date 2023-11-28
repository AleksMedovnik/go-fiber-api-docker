package products

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Delete Product
// @Tags Product
// @Description Delete Product
// @ID delete-product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Router /products/{id} [delete]
func (h handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&product)

	return c.SendStatus(fiber.StatusOK)
}