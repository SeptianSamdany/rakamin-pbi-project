package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rakamin-project/models"
	"rakamin-project/services"
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

// Upload Foto Produk
func (pc *ProdukController) UploadFotoProduk(c *fiber.Ctx) error {
	// Ambil ID Produk dari parameter
	produkID := c.Params("id")

	// Konversi ID Produk ke uint
	produkIDUint, err := strconv.ParseUint(produkID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	// Ambil file dari form-data
	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to get file"})
	}

	// Buat direktori uploads jika belum ada
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create upload directory"})
		}
	}

	// Generate nama file unik
	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), extension)
	filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// Simpan file ke server
	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Simpan informasi foto produk ke database
	fotoProduk := models.FotoProduk{
		IDProduk: uint(produkIDUint),
		Url:      fmt.Sprintf("http://localhost:8080/uploads/%s", filename),
	}

	// Simpan foto produk ke database
	if err := pc.ProdukService.SaveFotoProduk(fotoProduk); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save foto produk"})
	}

	// Update produk dengan URL foto yang baru
	produk, err := pc.ProdukService.GetProdukByID(uint(produkIDUint))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Menambahkan foto ke produk
	produk.FotoProduk = append(produk.FotoProduk, fotoProduk)

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "File uploaded successfully", "foto_url": fotoProduk.Url, "produk": produk})
}

// GetAllProduk handler
func (pc *ProdukController) GetAllProduk(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")    // default to page 1
	limitStr := c.Query("limit", "10") // default to 10 items per page
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

// Update Foto Produk (PUT)
func (pc *ProdukController) UpdateFotoProduk(c *fiber.Ctx) error {
	// Ambil ID produk dari parameter
	produkID := c.Params("id")

	// Konversi ID ke uint
	produkIDUint, err := strconv.ParseUint(produkID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	// Ambil file dari form-data
	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to get file"})
	}

	// Pastikan direktori uploads ada
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// Generate nama file unik
	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), extension)
	filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// Simpan file ke server
	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// URL yang akan disimpan
	fotoUrl := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	// Update foto produk di database
	err = pc.ProdukService.UpdateFotoProduk(uint(produkIDUint), fotoUrl)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Foto produk berhasil diperbarui", "foto_url": fotoUrl})
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

// Delete Foto Produk (DELETE)
func (pc *ProdukController) DeleteFotoProduk(c *fiber.Ctx) error {
	// Ambil ID produk dari parameter
	produkID := c.Params("id")

	// Konversi ID ke uint
	produkIDUint, err := strconv.ParseUint(produkID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	// Panggil service untuk menghapus foto produk
	err = pc.ProdukService.DeleteFotoProduk(uint(produkIDUint))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Foto produk berhasil dihapus"})
}
