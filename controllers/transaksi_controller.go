package controllers

import (
	"net/http"
	"rakamin-project/models"
	"rakamin-project/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TransaksiController struct {
	TransaksiService *services.TransaksiService
}

func NewTransaksiController(service *services.TransaksiService) *TransaksiController {
	return &TransaksiController{
		TransaksiService: service,
	}
}

// CreateTransaksi membuat transaksi baru
func (tc *TransaksiController) CreateTransaksi(c *fiber.Ctx) error {
	var input models.Transaksi
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi input
	if input.IDUser == 0 || input.AlamatPengirim == "" || input.HargaTotal == 0 || input.MethodBayar == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Buat transaksi
	transaksi, err := tc.TransaksiService.CreateTransaksi(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	// Return response
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Transaction created successfully", "data": transaksi})
}

// GetTransaksiByID mengambil transaksi berdasarkan ID
func (tc *TransaksiController) GetTransaksiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	transaksiID, err := strconv.Atoi(id) // Konversi string ke int
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	transaksi, err := tc.TransaksiService.GetTransaksiByID(uint(transaksiID)) // Konversi int ke uint
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Transaction not found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": transaksi})
}

// UpdateTransaksi memperbarui transaksi yang ada
func (tc *TransaksiController) UpdateTransaksi(c *fiber.Ctx) error {
	id := c.Params("id")
	transaksiID, err := strconv.Atoi(id) // Konversi string ke int
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var input models.Transaksi
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update transaksi
	transaksi, err := tc.TransaksiService.UpdateTransaksi(uint(transaksiID), input) // Konversi int ke uint
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update transaction"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Transaction updated successfully", "data": transaksi})
}

// DeleteTransaksi menghapus transaksi
func (tc *TransaksiController) DeleteTransaksi(c *fiber.Ctx) error {
	id := c.Params("id")
	transaksiID, err := strconv.Atoi(id) // Konversi string ke int
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := tc.TransaksiService.DeleteTransaksi(uint(transaksiID)); err != nil { // Konversi int ke uint
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete transaction"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Transaction deleted successfully"})
}

func (tc *TransaksiController) SaveDetailTransaksi(c *fiber.Ctx) error {
	var detail models.DetailTransaksi
	if err := c.BodyParser(&detail); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi detail transaksi
	if detail.IDTrx == 0 || detail.IDLogProduk == 0 || detail.Kuantitas == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Simpan detail transaksi
	detailTransaksi, err := tc.TransaksiService.SaveDetailTransaksi(detail)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save transaction details"})
	}

	// Return response
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Transaction detail saved successfully", "data": detailTransaksi})
}

func (tc *TransaksiController) CreateLogProduk(c *fiber.Ctx) error {
	var logProduk models.LogProduk
	if err := c.BodyParser(&logProduk); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi log produk
	if logProduk.IDProduk == 0 || logProduk.NamaProduk == "" || logProduk.HargaReseller == 0 || logProduk.HargaKonsumen == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Simpan log produk
	log, err := tc.TransaksiService.CreateLogProduk(logProduk)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create log for product"})
	}

	// Return response
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Product log created successfully", "data": log})
}

// Update Log Produk
func (tc *TransaksiController) UpdateLogProduk(c *fiber.Ctx) error {
	logID := c.Params("id") // Ambil ID log produk dari parameter

	// Konversi ID ke uint
	logIDUint, err := strconv.ParseUint(logID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid log ID"})
	}

	// Ambil data input dari request body
	var updatedLog models.LogProduk
	if err := c.BodyParser(&updatedLog); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service untuk update log produk
	logProduk, err := tc.TransaksiService.UpdateLogProduk(uint(logIDUint), updatedLog)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Log produk berhasil diperbarui", "data": logProduk})
}

// Delete Log Produk
func (tc *TransaksiController) DeleteLogProduk(c *fiber.Ctx) error {
	logID := c.Params("id") // Ambil ID log produk dari parameter

	// Konversi ID ke uint
	logIDUint, err := strconv.ParseUint(logID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid log ID"})
	}

	// Panggil service untuk delete log produk
	err = tc.TransaksiService.DeleteLogProduk(uint(logIDUint))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Log produk berhasil dihapus"})
}
