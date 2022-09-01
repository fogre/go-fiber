package trips

import (
	"go-fiber-app/package/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetTrips(c *fiber.Ctx) error {
	var trips []models.Trip

	if result := h.DB.Find(&trips); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&trips)
}
