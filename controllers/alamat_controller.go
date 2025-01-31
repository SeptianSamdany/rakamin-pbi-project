package controllers

import (
	"net/http"
	"rakamin-project/models"
	"rakamin-project/services"

	"github.com/gofiber/fiber/v2"
)

type AlamatController struct {
	AlamatService *services.AlamatService
}

// Create Alamat
func (ac *AlamatController) CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var alamatInput models.Alamat
	if err := c.BodyParser(&alamatInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	alamatInput.IDUser = userID
	alamat, err := ac.AlamatService.CreateAlamat(alamatInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(alamat)
}

// Get All Alamat
func (ac *AlamatController) GetAllAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	alamatList, err := ac.AlamatService.GetAllAlamat(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(alamatList)
}

// Get Alamat By ID
func (ac *AlamatController) GetAlamatByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	alamatID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid alamat ID"})
	}

	alamat, err := ac.AlamatService.GetAlamatByID(userID, uint(alamatID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(alamat)
}

// Update Alamat
func (ac *AlamatController) UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	alamatID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid alamat ID"})
	}

	var alamatInput models.Alamat
	if err := c.BodyParser(&alamatInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	updatedAlamat, err := ac.AlamatService.UpdateAlamat(userID, uint(alamatID), alamatInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(updatedAlamat)
}

// Delete Alamat
func (ac *AlamatController) DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	alamatID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid alamat ID"})
	}

	err = ac.AlamatService.DeleteAlamat(userID, uint(alamatID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Alamat deleted successfully"})
}
