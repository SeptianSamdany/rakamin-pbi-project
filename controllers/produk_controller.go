package controllers

import (
	"fmt"
	"net/http"
	"rakamin-project/models"
	"rakamin-project/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strings"
	"strconv"
)

type ProdukController struct {
	ProdukService *services.ProdukService
}

// Create Produk handler
func (pc *ProdukController) CreateProduk(c *fiber.Ctx) error {
	var produkInput models.Produk

	// Parse input JSON
	if err := c.BodyParser(&produkInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Call service to create produk
	produk, err := pc.ProdukService.CreateProduk(produkInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(produk)
}

// Upload Foto Produk handler
// Upload Foto Produk handler
func (pc *ProdukController) UploadFotoProduk(c *fiber.Ctx) error {
	produkID := c.Params("id")

	produkIDUint, err := strconv.Atoi(produkID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to get file"})
	}

	extension := file.Filename[strings.LastIndex(file.Filename, "."):]
	filename := fmt.Sprintf("%s%s", uuid.New().String(), extension)

	// Directory for uploads (make sure the folder exists)
	filepath := fmt.Sprintf("./uploads/%s", filename)

	// Save file to disk
	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Build URL for the uploaded file
	// Adjust the base URL depending on your setup (this is for example purposes)
	fileURL := fmt.Sprintf("http://localhost:3000/uploads/%s", filename)

	// Save the photo to the database
	fotoProduk := models.FotoProduk{
		IDProduk: uint(produkIDUint),
		Url:      fileURL,
	}

	if err := pc.ProdukService.SaveFotoProduk(fotoProduk); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save foto produk"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "File uploaded successfully"})
}

// GetAllProduk handler
// GetAllProduk handler with pagination
func (pc *ProdukController) GetAllProduk(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")      // default to page 1
	limitStr := c.Query("limit", "10")   // default to 10 items per page
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	produks, err := pc.ProdukService.GetAllProduk(page, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(produks)
}

// GetProdukByID handler
func (pc *ProdukController) GetProdukByID(c *fiber.Ctx) error {
	produkID := c.Params("id")

	// Convert ID ke uint
	id, err := strconv.Atoi(produkID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	produk, err := pc.ProdukService.GetProdukByID(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(produk)
}

// UpdateProduk handler
func (pc *ProdukController) UpdateProduk(c *fiber.Ctx) error {
	produkID := c.Params("id")

	// Convert ID ke uint
	id, err := strconv.Atoi(produkID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var produkInput models.Produk
	if err := c.BodyParser(&produkInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	produk, err := pc.ProdukService.UpdateProduk(uint(id), produkInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(produk)
}

// DeleteProduk handler
func (pc *ProdukController) DeleteProduk(c *fiber.Ctx) error {
	produkID := c.Params("id")

	// Convert ID ke uint
	id, err := strconv.Atoi(produkID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	err = pc.ProdukService.DeleteProduk(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
