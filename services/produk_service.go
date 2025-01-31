	package services

	import (
		"gorm.io/gorm"
		"rakamin-project/models"
		"errors"
		"os"
	)

	type ProdukService struct {
		DB *gorm.DB
	}

	func NewProdukService(db *gorm.DB) *ProdukService {
		return &ProdukService{DB: db}
	}
	
	// Create Produk - Menambahkan produk baru ke dalam database
	func (ps *ProdukService) CreateProduk(produkInput models.Produk) (*models.Produk, error) {
		// Validasi produk jika sudah ada berdasarkan slug
		var existingProduk models.Produk
		if err := ps.DB.Where("slug = ?", produkInput.Slug).First(&existingProduk).Error; err == nil {
			return nil, errors.New("produk dengan slug ini sudah ada")
		}
	
		// Simpan produk baru ke database
		if err := ps.DB.Create(&produkInput).Error; err != nil {
			return nil, err
		}
		return &produkInput, nil
	}
	
	// Save Foto Produk - Menyimpan foto produk yang baru
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
	
	// GetAllProduk - Mengambil semua produk dengan pagination
	func (ps *ProdukService) GetAllProduk(page, limit int) ([]models.Produk, error) {
		var produks []models.Produk
	
		// Preload untuk memuat relasi Toko, Kategori, dan FotoProduk
		if err := ps.DB.Preload("Toko").Preload("Kategori").Preload("FotoProduk").Limit(limit).Offset((page - 1) * limit).Find(&produks).Error; err != nil {
			return nil, err
		}
	
		return produks, nil
	}
	
	// GetProdukByID - Mengambil produk berdasarkan ID
	func (ps *ProdukService) GetProdukByID(id uint) (*models.Produk, error) {
		var produk models.Produk
	
		// Gunakan Preload untuk memuat relasi Toko, Kategori, dan FotoProduk
		if err := ps.DB.Preload("Toko").Preload("Kategori").Preload("FotoProduk").First(&produk, id).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}
	
		return &produk, nil
	}
	
	// UpdateProduk - Memperbarui data produk berdasarkan ID
	func (ps *ProdukService) UpdateProduk(id uint, updatedProduk models.Produk) (*models.Produk, error) {
		var produk models.Produk
	
		// Pastikan produk yang akan diupdate ada
		if err := ps.DB.First(&produk, id).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan")
		}
	
		// Update data produk
		if err := ps.DB.Model(&produk).Updates(updatedProduk).Error; err != nil {
			return nil, err
		}
	
		// Mengambil kembali produk yang sudah terupdate
		if err := ps.DB.Preload("Toko").Preload("Kategori").Preload("FotoProduk").First(&produk, id).Error; err != nil {
			return nil, errors.New("produk tidak ditemukan setelah update")
		}
	
		return &produk, nil
	}

	// Update Foto Produk
func (ps *ProdukService) UpdateFotoProduk(produkID uint, fotoUrl string) error {
	// Pastikan produk dengan ID yang sesuai ada
	var fotoProduk models.FotoProduk
	if err := ps.DB.Where("id_produk = ?", produkID).First(&fotoProduk).Error; err != nil {
		return errors.New("foto produk tidak ditemukan")
	}

	// Update URL foto
	fotoProduk.Url = fotoUrl

	// Simpan perubahan ke database
	if err := ps.DB.Save(&fotoProduk).Error; err != nil {
		return err
	}

	return nil
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

	// Delete Foto Produk
func (ps *ProdukService) DeleteFotoProduk(produkID uint) error {
	// Cari foto produk berdasarkan ID produk
	var fotoProduk models.FotoProduk
	if err := ps.DB.Where("id_produk = ?", produkID).First(&fotoProduk).Error; err != nil {
		return errors.New("foto produk tidak ditemukan")
	}

	// Hapus file gambar dari server
	if err := os.Remove(fotoProduk.Url); err != nil {
		return errors.New("gagal menghapus file foto produk")
	}

	// Hapus foto produk dari database
	if err := ps.DB.Delete(&fotoProduk).Error; err != nil {
		return err
	}

	return nil
}
