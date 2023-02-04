package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", DbDriver, DbUser, DbPassword, DbHost, DbPort, DbName)
	fmt.Println(DbURL)
	DB, err = gorm.Open(DbDriver, DbURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", DbDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", DbDriver)
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Project{})

}
