package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/Sc01100100/SaveCash-API/config"
    "github.com/Sc01100100/SaveCash-API/routes"
    "github.com/joho/godotenv"
)

func main() {
    if os.Getenv("APP_ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("Error loading .env file:", err)
        }
    }

    config.ConnectDB()
    defer config.Database.Close()

    log.Println("Database connection established.")

    app := fiber.New()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":8080"))
}