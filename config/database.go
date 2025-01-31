package config

import (
	"fmt"
	"log"
	"os"
	"rakamin-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB menginisialisasi koneksi database
func InitDB() *gorm.DB {
	// Konfigurasi database dari environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "rakamin_project"),
	)

	// Koneksi ke MySQL dengan logger agar lebih mudah debugging
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil
	}

	// Migrasi model dengan urutan yang benar
	err = db.AutoMigrate(
		&models.User{}, 
		&models.Toko{}, 
		&models.Alamat{}, 
		&models.Kategori{}, 
		&models.Produk{}, 
		&models.FotoProduk{}, 
		&models.LogProduk{}, 
		&models.Transaksi{}, 
		&models.DetailTransaksi{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful!")
	}

	return db
}

// getEnv mengambil nilai environment variable atau nilai default jika tidak tersedia
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}