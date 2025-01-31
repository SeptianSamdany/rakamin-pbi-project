package services

import (
	"errors"
	"rakamin-project/models"
	"gorm.io/gorm"
)

type TransaksiService struct {
	DB *gorm.DB
}

func NewTransaksiService(db *gorm.DB) *TransaksiService {
	return &TransaksiService{DB: db}
}

// SaveDetailTransaksi untuk menyimpan detail transaksi
func (ts *TransaksiService) SaveDetailTransaksi(detail models.DetailTransaksi) (*models.DetailTransaksi, error) {
	// Validasi produk yang terkait dengan detail transaksi
	var logProduk models.LogProduk
	if err := ts.DB.First(&logProduk, detail.IDLogProduk).Error; err != nil {
		return nil, errors.New("log produk tidak ditemukan")
	}

	// Validasi transaksi terkait dengan IDTrx
	var transaksi models.Transaksi
	if err := ts.DB.First(&transaksi, detail.IDTrx).Error; err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	// Simpan detail transaksi ke database
	if err := ts.DB.Create(&detail).Error; err != nil {
		return nil, err
	}

	return &detail, nil
}

// CreateLogProduk untuk mencatat produk dalam log
func (ts *TransaksiService) CreateLogProduk(logProduk models.LogProduk) (*models.LogProduk, error) {
	// Validasi produk berdasarkan IDProduk
	var produk models.Produk
	if err := ts.DB.First(&produk, logProduk.IDProduk).Error; err != nil {
		return nil, errors.New("produk tidak ditemukan")
	}

	// Simpan log produk
	if err := ts.DB.Create(&logProduk).Error; err != nil {
		return nil, err
	}

	return &logProduk, nil
}

// CreateTransaksi untuk membuat transaksi baru
func (ts *TransaksiService) CreateTransaksi(transaksi models.Transaksi) (*models.Transaksi, error) {
	// Validasi user dan alamat pengirim
	var user models.User
	if err := ts.DB.First(&user, transaksi.IDUser).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Simpan transaksi
	if err := ts.DB.Create(&transaksi).Error; err != nil {
		return nil, err
	}

	return &transaksi, nil
}

// GetTransaksiByID untuk mengambil transaksi berdasarkan ID
func (ts *TransaksiService) GetTransaksiByID(id uint) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	if err := ts.DB.Preload("DetailTransaksi").First(&transaksi, id).Error; err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}
	return &transaksi, nil
}

// UpdateTransaksi untuk memperbarui transaksi berdasarkan ID
func (ts *TransaksiService) UpdateTransaksi(id uint, transaksi models.Transaksi) (*models.Transaksi, error) {
	var existingTransaksi models.Transaksi
	if err := ts.DB.First(&existingTransaksi, id).Error; err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	// Update transaksi
	if err := ts.DB.Model(&existingTransaksi).Updates(transaksi).Error; err != nil {
		return nil, err
	}

	return &existingTransaksi, nil
}

// DeleteTransaksi untuk menghapus transaksi berdasarkan ID
func (ts *TransaksiService) DeleteTransaksi(id uint) error {
	var transaksi models.Transaksi
	if err := ts.DB.First(&transaksi, id).Error; err != nil {
		return errors.New("transaksi tidak ditemukan")
	}

	// Hapus transaksi
	if err := ts.DB.Delete(&transaksi).Error; err != nil {
		return err
	}

	return nil
}