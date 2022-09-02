package stations

import (
	"go-fiber-app/package/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetStations(c *fiber.Ctx) error {
	var stations []models.Station

	if result := h.DB.Find(&stations); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&stations)
}
