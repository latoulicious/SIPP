package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

// CreateSoalPayload is the struct used to match the JSON payload
// for creating a new soal. It contains the related user, mapel,
// kelas, jurusan, tahun ajar, bank soal, and dynamic field data.
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

// Field defines a dynamic field for a question, with a value and label.
// This is used in the CreateSoalPayload struct for dynamic question fields.

type Field struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type SoalHandler struct {
	SoalService *service.SoalService
}

// NewSoalHandler creates a new SoalHandler with the given SoalService.
// It is used to initialize a SoalHandler to handle soal related requests.

func NewSoalHandler(soalService *service.SoalService) *SoalHandler {
	return &SoalHandler{
		SoalService: soalService,
	}
}

// GetSoal retrieves all Soal records from the database along with their
// associated ItemSoal records.
//
// It accepts a Fiber context and returns a Fiber error.
//
// It calls the SoalService's GetSoal method to retrieve the soals and
// itemsoals from the database.
//
// It returns a 200 status code and success message with the soals data
// on success, or a 500 status code and error message on error.
// GetSoal retrieves all Soal records along with their associated ItemSoal records

func (handler *SoalHandler) GetSoal(c *fiber.Ctx) error {
	soals, err := handler.SoalService.GetSoal()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting soal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Soal retrieved successfully", "data": soals})
}

// GetSoalByID retrieves a Soal record by ID along with its associated ItemSoal records.
// It takes a Soal ID as a parameter, parses it into a UUID, retrieves the Soal record by that ID from the SoalService, and returns it in the response.
// Returns a 400 status code if the ID is invalid.
// Returns a 500 status code if there is an error retrieving the Soal.
// Otherwise returns a 200 status code with the Soal record in the response body.

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

// CreateSoal handles creating a new Soal and associated ItemSoal record.
// It parses the request body into a CreateSoalPayload, constructs a Soal
// and ItemSoal model instance from the payload, calls the service layer
// to persist them, and returns the created ItemSoal in the response.

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

// UpdateSoal updates an existing soal record in the database.
// It takes the soal ID as a parameter, parses the updated soal data from the request body,
// calls the SoalService to update the soal, and returns the updated soal.
// Returns 400 if the soal ID is invalid.
// Returns 500 if there is an error parsing the request body or updating the soal.
// Returns 404 if the soal is not found after updating.

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

// DeleteSoal deletes a soal by ID.
// It parses the soal ID from the request parameters, calls the SoalService
// to delete the soal, and returns a JSON response indicating success or failure.

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
