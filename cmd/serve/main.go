package main

import (
	"blog/database"
	"blog/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()

	err := godotenv.Load()

	if err != nil {
		log.Println("File env tidak ditemukan")
	}

	port := os.Getenv("APP_PORT")

	app := fiber.New()

	routes.Routes(app)

	app.Listen(":" + port)
}
