package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type TahunHandler struct {
	TahunService *service.TahunService
}

func NewTahunHandler(tahunService *service.TahunService) *TahunHandler {
	return &TahunHandler{
		TahunService: tahunService,
	}
}

func (handler *TahunHandler) GetTahun(c *fiber.Ctx) error {
	tahun, err := handler.TahunService.GetTahun()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting users", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tahun retrieved successfully", "data": tahun})
}

func (handler *TahunHandler) GetTahunByID(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun retrieved successfully", "data": tahunID})
}

func (handler *TahunHandler) CreateTahun(c *fiber.Ctx) error {
	tahun := new(model.TahunAjar)

	err := c.BodyParser(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.TahunService.CreateTahun(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun created", "data": tahun})
}

func (handler *TahunHandler) UpdateTahun(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	tahun, err := handler.TahunService.GetTahunByID(tahunID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Tahun not found", "data": nil})
	}

	err = c.BodyParser(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.TahunService.UpdateTahun(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun updated", "data": tahun})
}

func (handler *TahunHandler) DeleteTahun(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	err = handler.TahunService.DeleteTahun(tahunID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun deleted"})
}
