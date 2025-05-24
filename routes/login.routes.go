package routes

import (
	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var signUp models.User

	if err := c.BodyParser(&signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(signUp.Password), bcrypt.DefaultCost)
	signUp.Password = string(hashedPassword)

	if err := db.DB.Create(&signUp).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(signUp)

}

func SignIn(c *fiber.Ctx) error {
	var input models.SignIn
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	user.Password = "********"

	return c.JSON(user)
}
