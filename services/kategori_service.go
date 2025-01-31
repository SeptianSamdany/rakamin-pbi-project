package services

import (
	"errors"
	"rakamin-project/models"

	"gorm.io/gorm"
)

type KategoriService struct {
	DB *gorm.DB
}

func NewKategoriService(db *gorm.DB) *KategoriService {
	return &KategoriService{DB: db}
}

// Create kategori (hanya untuk admin)
func (ks *KategoriService) CreateKategori(userID uint, namaKategori string) (*models.Kategori, error) {
	// Cek apakah user adalah admin
	var user models.User
	if err := ks.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	if !user.IsAdmin {
		return nil, errors.New("hanya admin yang bisa menambahkan kategori")
	}

	kategori := models.Kategori{
		NamaKategori: namaKategori,
	}
	if err := ks.DB.Create(&kategori).Error; err != nil {
		return nil, err
	}

	return &kategori, nil
}

// Get all kategori
func (ks *KategoriService) GetAllKategori() ([]models.Kategori, error) {
	var kategori []models.Kategori
	if err := ks.DB.Find(&kategori).Error; err != nil {
		return nil, err
	}
	return kategori, nil
}

// Update kategori (hanya untuk admin)
func (ks *KategoriService) UpdateKategori(userID uint, kategoriID uint, namaKategori string) (*models.Kategori, error) {
	// Cek apakah user adalah admin
	var user models.User
	if err := ks.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	if !user.IsAdmin {
		return nil, errors.New("hanya admin yang bisa mengupdate kategori")
	}

	// Cari kategori
	var kategori models.Kategori
	if err := ks.DB.First(&kategori, kategoriID).Error; err != nil {
		return nil, errors.New("kategori tidak ditemukan")
	}

	// Update kategori
	kategori.NamaKategori = namaKategori
	if err := ks.DB.Save(&kategori).Error; err != nil {
		return nil, err
	}

	return &kategori, nil
}

// Delete kategori (hanya untuk admin)
func (ks *KategoriService) DeleteKategori(userID uint, kategoriID uint) error {
	// Cek apakah user adalah admin
	var user models.User
	if err := ks.DB.First(&user, userID).Error; err != nil {
		return errors.New("user tidak ditemukan")
	}
	if !user.IsAdmin {
		return errors.New("hanya admin yang bisa menghapus kategori")
	}

	// Hapus kategori
	if err := ks.DB.Delete(&models.Kategori{}, kategoriID).Error; err != nil {
		return err
	}

	return nil
}
