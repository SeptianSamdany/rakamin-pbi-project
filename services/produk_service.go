	package services

	import (
		"gorm.io/gorm"
		"rakamin-project/models"
		"errors"
	)

	type ProdukService struct {
		DB *gorm.DB
	}

	func NewProdukService(db *gorm.DB) *ProdukService {
		return &ProdukService{DB: db}
	}

	// Create Produk
	func (ps *ProdukService) CreateProduk(produkInput models.Produk) (*models.Produk, error) {
		if err := ps.DB.Create(&produkInput).Error; err != nil {
			return nil, err
		}
		return &produkInput, nil
	}

	// Save Foto Produk
	func (ps *ProdukService) SaveFotoProduk(fotoProduk models.FotoProduk) error {
		// Pastikan produk dengan ID yang sesuai ada
		var produk models.Produk
		if err := ps.DB.First(&produk, fotoProduk.IDProduk).Error; err != nil {
			return errors.New("produk tidak ditemukan")
		}

		// Simpan foto produk ke database
		if err := ps.DB.Create(&fotoProduk).Error; err != nil {
			return err
		}

		return nil
	}

	// GetAllProduk - Mengambil semua produk dari database
	func (ps *ProdukService) GetAllProduk(page, limit int) ([]models.Produk, error) {
		var produks []models.Produk
		if err := ps.DB.Preload("Toko").Preload("Kategori").Limit(limit).Offset((page - 1) * limit).Find(&produks).Error; err != nil {
			return nil, err
		}
		return produks, nil
	}

	// GetProdukByID - Mengambil produk berdasarkan ID
	func (ps *ProdukService) GetProdukByID(id uint) (*models.Produk, error) {
		var produk models.Produk
		// Gunakan Preload untuk memuat relasi Toko dan Kategori
		if err := ps.DB.Preload("Toko").Preload("Kategori").First(&produk, id).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}
		return &produk, nil
	}

	// UpdateProduk - Memperbarui data produk berdasarkan ID
	func (ps *ProdukService) UpdateProduk(id uint, updatedProduk models.Produk) (*models.Produk, error) {
		var produk models.Produk
		if err := ps.DB.First(&produk, id).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}

		// Update data produk
		if err := ps.DB.Model(&produk).Updates(updatedProduk).Error; err != nil {
			return nil, err
		}

		return &produk, nil
	}

	// DeleteProduk - Menghapus produk berdasarkan ID
	func (ps *ProdukService) DeleteProduk(id uint) error {
		var produk models.Produk
		if err := ps.DB.First(&produk, id).Error; err != nil {
			return errors.New("produk tidak ditemukan")
		}

		// Hapus produk dari database
		if err := ps.DB.Delete(&produk).Error; err != nil {
			return err
		}

		return nil
	}
