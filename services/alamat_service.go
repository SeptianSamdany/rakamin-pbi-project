package services

import (
	"errors"

	"gorm.io/gorm"
	"rakamin-project/models"
)

type AlamatService struct {
	DB *gorm.DB
}

func NewAlamatService(db *gorm.DB) *AlamatService {
	return &AlamatService{DB: db}
}

// Create Alamat
func (as *AlamatService) CreateAlamat(alamatInput models.Alamat) (*models.Alamat, error) {
	if err := as.DB.Create(&alamatInput).Error; err != nil {
		return nil, err
	}
	return &alamatInput, nil
}

// Get All Alamat by UserID
func (as *AlamatService) GetAllAlamat(userID uint) ([]models.Alamat, error) {
	var alamatList []models.Alamat
	if err := as.DB.Where("id_user = ?", userID).Find(&alamatList).Error; err != nil {
		return nil, err
	}
	return alamatList, nil
}

// Get Alamat by ID
func (as *AlamatService) GetAlamatByID(userID, alamatID uint) (*models.Alamat, error) {
	var alamat models.Alamat
	if err := as.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return nil, errors.New("alamat tidak ditemukan atau tidak memiliki akses")
	}
	return &alamat, nil
}

// Update Alamat
func (as *AlamatService) UpdateAlamat(userID, alamatID uint, alamatInput models.Alamat) (*models.Alamat, error) {
	var existingAlamat models.Alamat

	// Cek apakah alamat ada dan milik user yang bersangkutan
	if err := as.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&existingAlamat).Error; err != nil {
		return nil, errors.New("alamat tidak ditemukan atau tidak memiliki akses")
	}

	// Update data alamat
	if err := as.DB.Model(&existingAlamat).Updates(alamatInput).Error; err != nil {
		return nil, err
	}

	return &existingAlamat, nil
}

// Delete Alamat
func (as *AlamatService) DeleteAlamat(userID, alamatID uint) error {
	var alamat models.Alamat

	// Cek apakah alamat ada dan milik user yang bersangkutan
	if err := as.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return errors.New("alamat tidak ditemukan atau tidak memiliki akses")
	}

	// Hapus alamat
	if err := as.DB.Delete(&alamat).Error; err != nil {
		return err
	}

	return nil
}
