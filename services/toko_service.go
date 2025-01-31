package services

import (
	"errors"
	"rakamin-project/models"
	"gorm.io/gorm"
)

// TokoService mengelola logika bisnis terkait Toko
type TokoService struct {
	DB *gorm.DB
}

func NewTokoService(db *gorm.DB) *TokoService {
	return &TokoService{DB: db}
}

// CreateToko membuat Toko baru untuk user
func (ts *TokoService) CreateToko(userID uint, namaToko string) (*models.Toko, error) {
	// Cek apakah user sudah memiliki toko
	var toko models.Toko
	if err := ts.DB.Where("id_user = ?", userID).First(&toko).Error; err == nil {
		return nil, errors.New("user sudah memiliki toko")
	}

	// Buat Toko baru
	newToko := models.Toko{
		IDUser:   userID,
		NamaToko: namaToko,
		UrlToko:  "url-" + namaToko, // Contoh, bisa dibuat dinamis
	}

	if err := ts.DB.Create(&newToko).Error; err != nil {
		return nil, err
	}

	return &newToko, nil
}

// UpdateToko memperbarui data Toko user
func (ts *TokoService) UpdateToko(userID uint, tokoID uint, namaToko string) (*models.Toko, error) {
	var toko models.Toko
	if err := ts.DB.Where("id_user = ? AND id = ?", userID, tokoID).First(&toko).Error; err != nil {
		return nil, errors.New("toko tidak ditemukan atau user tidak memiliki akses")
	}

	toko.NamaToko = namaToko
	if err := ts.DB.Save(&toko).Error; err != nil {
		return nil, err
	}

	return &toko, nil
}

func (ts *TokoService) GetTokoByUserID(userID uint) (*models.Toko, error) {
    var toko models.Toko
    
    if err := ts.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
        return nil, errors.New("toko tidak ditemukan")
    }

    return &toko, nil
}

// DeleteToko menghapus Toko user
func (ts *TokoService) DeleteToko(userID uint, tokoID uint) error {
	var toko models.Toko
	if err := ts.DB.Where("id_user = ? AND id = ?", userID, tokoID).First(&toko).Error; err != nil {
		return errors.New("toko tidak ditemukan atau user tidak memiliki akses")
	}

	if err := ts.DB.Delete(&toko).Error; err != nil {
		return err
	}

	return nil
}
