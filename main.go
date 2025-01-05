package main

import (
	"log"
	"os"

	"github.com/Sc01100100/SaveCash-API/config"
	"github.com/Sc01100100/SaveCash-API/middlewares"
	"github.com/Sc01100100/SaveCash-API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    config.ConnectDB()
    defer config.Database.Close()

    log.Println("Database connection established.")

    app := fiber.New()

    app.Use(cors.New(middlewares.Cors))
   
    routes.SetupRoutes(app)
   
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }
    log.Fatal(app.Listen(":" + port))
}