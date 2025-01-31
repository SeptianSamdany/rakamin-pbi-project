package controllers

import (
	"net/http"
	"rakamin-project/services"

	"github.com/gofiber/fiber/v2"
)

type KategoriController struct {
	KategoriService *services.KategoriService
}

// Create kategori (hanya admin)
func (kc *KategoriController) CreateKategori(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Dapatkan user_id dari middleware JWT

	var input struct {
		NamaKategori string `json:"nama_kategori"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	kategori, err := kc.KategoriService.CreateKategori(userID, input.NamaKategori)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(kategori)
}

// Get all kategori
func (kc *KategoriController) GetAllKategori(c *fiber.Ctx) error {
	kategori, err := kc.KategoriService.GetAllKategori()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(kategori)
}

// Update kategori (hanya admin)
func (kc *KategoriController) UpdateKategori(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		NamaKategori string `json:"nama_kategori"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	kategoriID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid kategori ID"})
	}

	kategori, err := kc.KategoriService.UpdateKategori(userID, uint(kategoriID), input.NamaKategori)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(kategori)
}

// Delete kategori (hanya admin)
func (kc *KategoriController) DeleteKategori(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	kategoriID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid kategori ID"})
	}

	err = kc.KategoriService.DeleteKategori(userID, uint(kategoriID))
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Kategori deleted successfully"})
}
