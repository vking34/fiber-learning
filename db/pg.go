package db

import (
	"fmt"
	"log"
	"os"

	"github.com/vking34/fiber-messenger/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB postgre DB
var DB *gorm.DB

// ConnectDB connect to postgre
func ConnectDB() {
	host := os.Getenv("POSTGRE_HOST")
	port := os.Getenv("POSTGRE_PORT")
	user := os.Getenv("POSTGRE_USER")
	pass := os.Getenv("POSTGRE_PASSWORD")
	db := os.Getenv("POSTGRE_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	log.Println("dsn:", dsn)

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect db")
	}

	log.Println("Connected to Postgre")
	DB.AutoMigrate(&models.User{})
}
