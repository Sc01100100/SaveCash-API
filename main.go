package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/Sc01100100/SaveCash-API/config"
	"github.com/Sc01100100/SaveCash-API/routes"
)

func main() {
    if os.Getenv("HEROKU") != "true" {
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }
    }

    config.ConnectDB()
    defer config.Database.Close()

    log.Println("Database connection established.")

    app := fiber.New()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":8080"))
}