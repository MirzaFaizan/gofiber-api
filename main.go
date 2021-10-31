package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mirzafaizan/gofiber-api/database"
	"github.com/mirzafaizan/gofiber-api/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRouter(app)
	log.Fatal(app.Listen(":8000"))

	defer database.DB.Close()
}
