package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

// Define the CreateSoalPayload struct to match the JSON payload structure
type CreateSoalPayload struct {
	User          model.Users     `json:"User"`
	Mapel         model.Mapel     `json:"Mapel"`
	Kelas         model.Kelas     `json:"Kelas"`
	Jurusan       model.Jurusan   `json:"Jurusan"`
	TahunAjar     model.TahunAjar `json:"TahunAjar"`
	BankSoal      model.BankSoal  `json:"BankSoal"`
	TipeSoal      string          `json:"TipeSoal"`
	Hari          string          `json:"Hari"`
	Tanggal       string          `json:"Tanggal"`
	Waktu         string          `json:"Waktu"`
	UserID        uuid.UUID       `json:"UserID"`
	MapelID       uuid.UUID       `json:"MapelID"`
	KelasID       uuid.UUID       `json:"KelasID"`
	JurusanID     uuid.UUID       `json:"JurusanID"`
	TahunAjarID   uuid.UUID       `json:"TahunAjarID"`
	BankSoalID    uuid.UUID       `json:"BankSoalID"`
	DynamicFields []Field         `json:"DynamicFields"`
}

type Field struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type SoalHandler struct {
	SoalService *service.SoalService
}

// soal handler
func NewSoalHandler(soalService *service.SoalService) *SoalHandler {
	return &SoalHandler{
		SoalService: soalService,
	}
}

// GetSoal retrieves all Soal records along with their associated ItemSoal records
func (handler *SoalHandler) GetSoal(c *fiber.Ctx) error {
	soals, err := handler.SoalService.GetSoal()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting soal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Soal retrieved successfully", "data": soals})
}

// GetSoalByID retrieves a Soal by ID along with its associated ItemSoal records
func (handler *SoalHandler) GetSoalByID(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	soal, err := handler.SoalService.GetSoalByID(soalID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting soal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Soal retrieved successfully", "data": soal})
}

func (handler *SoalHandler) CreateSoal(c *fiber.Ctx) error {
	// Parse the request payload into a CreateSoalPayload instance
	var payload CreateSoalPayload
	if err := c.BodyParser(&payload); err != nil {
		// ... your error handling for bad requests
		return err
	}

	// Construct the Soal instance from the payload
	soal := &model.Soal{
		ID:        uuid.New(),
		UserID:    payload.UserID,
		MapelID:   payload.MapelID,
		KelasID:   payload.KelasID,
		JurusanID: payload.JurusanID,
		TipeSoal:  payload.TipeSoal,
		Hari:      payload.Hari,
		Tanggal:   payload.Tanggal,
		Waktu:     payload.Waktu,
		// ... Map any other necessary fields
	}

	// Construct the dynamic fields map for ItemSoal
	dynamicFieldsMap := make(map[string]interface{})
	for _, field := range payload.DynamicFields {
		dynamicFieldsMap[field.Label] = field.Value
	}

	// Construct the ItemSoal instance from the payload
	itemSoal := &model.ItemSoal{
		ID:            uuid.New(),
		BankSoalID:    payload.BankSoalID,
		DynamicFields: dynamicFieldsMap,
		// ... Map any other necessary fields
	}

	// Call the service function to create the Soal and ItemSoal records
	if err := handler.SoalService.CreateSoal(soal, itemSoal); err != nil {
		// ... your error handling for errors coming from your service
		return err
	}

	// Construct the response body with the created ItemSoal
	responseBody := fiber.Map{
		"message": "Soal and ItemSoal created successfully",
		"item_soal": fiber.Map{
			"id":             itemSoal.ID,
			"bank_soal_id":   itemSoal.BankSoalID,
			"dynamic_fields": itemSoal.DynamicFields,
		},
	}

	// Log the successful creation
	log.Printf("Soal and ItemSoal created successfully. ID: %s\n", itemSoal.ID)

	// Return the response with the created ItemSoal
	return c.Status(fiber.StatusCreated).JSON(responseBody)
}

// UpdateSoal updates an existing soal
func (handler *SoalHandler) UpdateSoal(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	soal := new(model.Soal)
	soal.ID = soalID

	// Parse the JSON body into the existing Soal struct
	err = c.BodyParser(soal)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Soal in the database
	err = handler.SoalService.UpdateSoal(soal)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update soal", "data": err})
	}

	// Retrieve the existing Soal by ID
	soal, err = handler.SoalService.GetSoalByID(soalID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Soal not found", "data": nil})
	}

	// Return the updated Soal
	return c.JSON(fiber.Map{"status": "success", "message": "Soal updated", "data": soal})
}

// DeleteSoal deletes a soal by ID
func (handler *SoalHandler) DeleteSoal(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Soal from the database
	err = handler.SoalService.DeleteSoal(soalID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete soal", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Soal deleted", "data": nil})
}
