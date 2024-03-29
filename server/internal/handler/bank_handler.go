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

// NewBankHandler creates a new BankHandler with the given BankService.
func NewBankHandler(bankService *service.BankService) *BankHandler {
	return &BankHandler{
		BankService: bankService,
	}
}

// GetBank retrieves the bank details from the database.
// It calls the BankService's GetBank method to get the bank,
// and returns the bank data in the response if successful,
// or returns a 500 error response if there is an error.

func (handler *BankHandler) GetBank(c *fiber.Ctx) error {
	bank, err := handler.BankService.GetBank()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting bank", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "BankSoal retrieved successfully", "data": bank})
}

// GetBankByID retrieves a bank by its ID.
// It parses the ID from the route parameter "id", calls the BankService
// to get the bank by that ID, and returns the bank in the response if found.
// If the ID is invalid or the bank is not found, it returns a 400 or 500 error.

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

// CreateBank handles creating a new bank in the database.
// It parses the request body into a BankSoal model, calls the BankService's
// CreateBank method to insert into the database, and returns the created bank
// in the response if successful, or returns a 500 error if there is an error.
// The function logs helpful debugging info.

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

// UpdateBank updates the bank with the given ID by parsing the ID
// from the route parameter "id", parsing the updated bank data from
// the request body, calling the BankService to update the bank,
// and returning the updated bank in the response. It returns a 400
// error for an invalid ID, 500 error for parsing errors, and 404
// error if the bank is not found after updating.

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

// DeleteBank deletes a BankSoal record from the database by ID.
// It parses the bank ID from the URL parameter, calls the BankService
// to delete the bank, and returns a JSON response indicating success or failure.
// Returns a 400 if the ID is invalid, 404 if the bank is not found,
// or 500 if there is a database error.

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
