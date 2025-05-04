package routes

import (
	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	var users []models.SignUp

	if err := db.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(users)
}

func FindUser(c *fiber.Ctx) error {
	var user models.SignUp
	id := c.Params("id")

	if err := db.DB.First(&user, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}