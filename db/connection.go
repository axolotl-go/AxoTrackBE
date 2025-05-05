package db

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DSN = os.Getenv("DB_DSN")
var DB *gorm.DB

func Dbconnection() {
	var error error
	DB, error = gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connected successfully")
	}

}
