package stations

import (
	"go-fiber-app/package/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetTrips(c *fiber.Ctx) error {
	var stations []models.Trip

	if result := h.DB.Find(&stations); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&stations)
}
