package main

import (
	"log"
	"os"

	"github.com/axolotl-go/axoTrackBackEnd/db"
	"github.com/axolotl-go/axoTrackBackEnd/models"
	"github.com/axolotl-go/axoTrackBackEnd/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")

	db.Dbconnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.SignIn{})
	db.DB.AutoMigrate(models.Transaction{})

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", routes.HelloApp)
	app.Post("/signIn", routes.SignIn)
	app.Post("/signUp", routes.SignUp)

	// Users route
	app.Get("/users", routes.Users)
	app.Get("/user/:id", routes.FindUser)

	// Transactions route
	app.Post("/transaction", routes.Transaction)
	app.Post("/viewTransaction", routes.ViewTransaction)

	app.Listen(":" + DB_HOST)
}
