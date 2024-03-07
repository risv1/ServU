package db

import (
	"ServU/src/models"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	dbUser    string
	dbPassword string
	database  string
)

func init() {
	loadenv := godotenv.Load("../.env")
	if loadenv != nil {
		log.Println("Error loading env")
	}

	database = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PWD")
}

func Connect() {
	connection, err := gorm.Open(mysql.Open(dbUser + ":" + dbPassword + "@/" + database), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}