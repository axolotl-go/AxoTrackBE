package routes

import (
	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	db.DB.Model(&users).Association("Transactions").Find(&users)

	return c.JSON(users)
}

func FindUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")

	if err := db.DB.First(&user, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	db.DB.Model(&user).Association("Transactions").Find(&user.Transactions)

	return c.JSON(user)
}
