package routes

import (
	"net/http"

	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/gofiber/fiber/v2"
)

func Transaction(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})

	}

	return c.JSON(transaction)
}

func ViewTransaction(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := db.DB.Where("user_id = ?", transaction.UserID).Find(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve transaction",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"type":        transaction.TransactionType,
		"amount":      transaction.Amount,
		"description": transaction.Description,
		"category":    transaction.Category,
		"date":        transaction.Date,
	})

}
