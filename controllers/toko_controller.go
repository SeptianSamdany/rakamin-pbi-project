package controllers

import (
	"net/http"
	"rakamin-project/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TokoController struct {
	TokoService *services.TokoService
}

// CreateToko handler
func (tc *TokoController) CreateToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Dapatkan user_id dari middleware JWT

	var input struct {
		NamaToko string `json:"nama_toko"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service
	toko, err := tc.TokoService.CreateToko(userID, input.NamaToko)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(toko)
}

// UpdateToko handler
func (tc *TokoController) UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Dapatkan user_id dari middleware JWT

	// Ambil ID Toko dari URL parameter dan konversi menjadi uint
	tokoID, err := strconv.ParseUint(c.Params("id"), 10, 32) // Konversi string ke uint
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid toko ID"})
	}

	var input struct {
		NamaToko string `json:"nama_toko"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service
	toko, err := tc.TokoService.UpdateToko(userID, uint(tokoID), input.NamaToko)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(toko)
}

// GetToko handler
func (tc *TokoController) GetToko(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uint) // Ambil user ID dari JWT

    // Panggil service untuk mendapatkan data toko
    toko, err := tc.TokoService.GetTokoByUserID(userID)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(toko)
}

// DeleteToko handler
func (tc *TokoController) DeleteToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Dapatkan user_id dari middleware JWT

	// Ambil ID Toko dari URL parameter dan konversi menjadi uint
	tokoID, err := strconv.ParseUint(c.Params("id"), 10, 32) // Konversi string ke uint
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid toko ID"})
	}

	// Panggil service
	err = tc.TokoService.DeleteToko(userID, uint(tokoID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Toko berhasil dihapus"})
}
