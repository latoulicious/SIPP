package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type BankHandler struct {
	BankService *service.BankService
}

func NewBankHandler(bankService *service.BankService) *BankHandler {
	return &BankHandler{
		BankService: bankService,
	}
}

func (handler *BankHandler) GetBank(c *fiber.Ctx) error {
	bank, err := handler.BankService.GetBank()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting bank", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "BankSoal retrieved successfully", "data": bank})
}

func (handler *BankHandler) GetBankByID(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	bank, err := handler.BankService.GetBankByID(bankID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting bank", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "BankSoal retrieved successfully", "data": bank})
}

func (handler *BankHandler) CreateBank(c *fiber.Ctx) error {

	bank := new(model.BankSoal)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", bank)
	log.Printf("Received BankSoal object: %+v\n", bank)

	err := c.BodyParser(bank)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.BankService.CreateBank(bank)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating bank: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create bank", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "BankSoal created", "data": bank})
}

func (handler *BankHandler) UpdateBank(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	bank := new(model.BankSoal)
	bank.ID = bankID

	// Parse the JSON body into the existing BankSoal struct
	err = c.BodyParser(bank)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated BankSoal: %+v\n", bank)
	log.Print(bank.UserID)

	// Update the BankSoal in the database
	err = handler.BankService.UpdateBank(bank)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update bank", "data": err})
	}

	// Retrieve the existing BankSoal by ID
	bank, err = handler.BankService.GetBankByID(bankID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "BankSoal not found", "data": nil})
	}

	// Return the updated BankSoal
	return c.JSON(fiber.Map{"status": "success", "message": "BankSoal updated", "data": bank})
}

func (handler *BankHandler) DeleteBank(c *fiber.Ctx) error {
	bankID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the BankSoal from the database
	err = handler.BankService.DeleteBank(bankID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete bank", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "BankSoal deleted", "data": nil})
}
