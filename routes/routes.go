package routes

import (
	"github.com/Sc01100100/SaveCash-API/controllers"
	"github.com/Sc01100100/SaveCash-API/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/savecash")

	api.Post("/register", controllers.InsertUser)
	api.Post("/login", controllers.LoginUser)
	api.Post("/logout", controllers.LogoutUser)

	protected := api.Group("/", middlewares.AuthMiddleware())

	protected.Post("/transactions", controllers.CreateTransactionHandler)
	protected.Get("/transactions", controllers.GetTransactionsHandler)
	protected.Delete("/transactions/:id", controllers.DeleteTransactionHandler)

	protected.Post("/incomes", controllers.CreateIncomeHandler)
	protected.Get("/incomes", controllers.GetIncomesHandler)
	protected.Delete("/incomes/:id", controllers.DeleteIncomeHandler)

	protected.Get("/user/info", controllers.GetUserInfo)

	admin := protected.Group("/admin") 
	admin.Use(middlewares.AdminMiddleware()) 
	admin.Get("/users", controllers.GetAllUser) 
}