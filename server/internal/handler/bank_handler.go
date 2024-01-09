package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type BankHandler struct {
	bankService *service.BankService
}

func NewBankHandler(bankService *service.BankService) *BankHandler {
	return &BankHandler{
		bankService: bankService,
	}
}

func (handler *BankHandler) GetBank(c *fiber.Ctx) error {
	bank, err := handler.bankService.GetBank()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting bank", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Bank retrieved successfully", "data": bank})
}

func (handler *BankHandler) GetBankByID(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	bank, err := handler.bankService.GetBankByID(bankID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting bank", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Bank retrieved successfully", "data": bank})
}

func (handler *BankHandler) CreateBank(c *fiber.Ctx) error {
	var bank model.BankSoal

	// Log Request Body for debugging purposes
	log.Printf("Request Body: %+v\n", bank)

	err := c.BodyParser(&bank)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.bankService.CreateBank(&bank)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating bank: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create bank", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Bank created", "data": bank})
}

func (handler *BankHandler) UpdateBank(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Retrieve the existing Bank by ID
	bank, err := handler.bankService.GetBankByID(bankID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Bank not found", "data": nil})
	}

	// Parse the JSON body into the existing Bank struct
	err = c.BodyParser(bank)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Bank in the database
	err = handler.bankService.UpdateBank(bank)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update bank", "data": err})
	}

	// Return the updated Bank
	return c.JSON(fiber.Map{"status": "success", "message": "Bank updated", "data": bank})
}

func (handler *BankHandler) DeleteBank(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Bank from the database
	err = handler.bankService.DeleteBank(bankID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete bank", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Bank deleted", "data": nil})
}
