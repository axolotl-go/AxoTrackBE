package routes

import (
	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/gofiber/fiber/v2"
)

func CreateClient(c *fiber.Ctx) error {
	var loan models.Loan

	if err := c.BodyParser(&loan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := db.DB.Create(&loan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create loan",
		})
	}

	return c.JSON(loan)

}
