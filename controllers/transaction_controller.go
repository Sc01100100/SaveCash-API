package controllers

import (
	"strconv"
	"log"

	"github.com/Sc01100100/SaveCash-API/models"
	"github.com/Sc01100100/SaveCash-API/module"
	"github.com/gofiber/fiber/v2"
)

func CreateTransactionHandler(c *fiber.Ctx) error {
	var transaction models.Transaction

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction data",
		})
	}

	if transaction.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Amount must be greater than zero",
		})
	}
	if transaction.Category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Category cannot be empty",
		})
	}

	newTransaction, err := module.CreateTransaction(transaction.UserID, transaction.Amount, transaction.Category, transaction.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

	return c.JSON(fiber.Map{
		"status":      "success",
		"transaction": newTransaction,
	})
}

func DeleteTransactionHandler(c *fiber.Ctx) error {
	transactionID, err := strconv.Atoi(c.Params("transactionID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction ID",
		})
	}

	if err := module.DeleteTransaction(transactionID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete transaction",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transaction deleted successfully",
	})
}

func CreateIncomeHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		log.Println("UserID is missing in context")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "UserID is missing in context",
		})
	}

	intUserID, ok := userID.(int)
	if !ok || intUserID == 0 {
		log.Printf("Invalid UserID from context: %v\n", userID)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid UserID format",
		})
	}

	log.Printf("Creating income for UserID: %d\n", intUserID)

	var income models.Income
	if err := c.BodyParser(&income); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid income data",
		})
	}

	if income.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Amount must be greater than zero",
		})
	}
	if income.Source == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Source cannot be empty",
		})
	}

	newIncome, err := module.CreateIncome(intUserID, income.Amount, income.Source)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create income",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"income":  newIncome,
	})
}

func DeleteIncomeHandler(c *fiber.Ctx) error {
	incomeID, err := strconv.Atoi(c.Params("incomeID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid income ID",
		})
	}

	if err := module.DeleteIncome(incomeID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete income",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Income deleted successfully",
	})
}
