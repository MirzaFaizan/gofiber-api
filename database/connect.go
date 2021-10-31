package database

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/mirzafaizan/gofiber-api/config"
	"github.com/mirzafaizan/gofiber-api/models"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASS"), config.Config("DB_NAME")))

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migration complete")
}
