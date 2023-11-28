package products

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

// @Summary Put Product
// @Tags Product
// @Description Put Product
// @ID put-product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Param input body UpdateProductRequestBody true "data for product updates"
// @Router /products/{id} [put]
func (h handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	h.DB.Save(&product)

	return c.Status(fiber.StatusOK).JSON(&product)
}