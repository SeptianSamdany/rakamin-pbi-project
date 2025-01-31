package models

import (
	"time"
)

type FotoProduk struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk  uint      `gorm:"not null" json:"id_produk"` 
	Url       string    `gorm:"not null" json:"url"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Produk struct {
	ID             uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk     string       `gorm:"not null" json:"nama_produk"`
	Slug           string       `gorm:"unique;not null" json:"slug"`
	HargaReseller  int          `gorm:"not null" json:"harga_reseller"`
	HargaKonsumen  int          `gorm:"not null" json:"harga_konsumen"`
	Stok           int          `gorm:"not null" json:"stok"`
	Deskripsi      string       `json:"deskripsi"`
	IDToko         uint         `gorm:"not null" json:"id_toko"`
	Toko           Toko         `gorm:"foreignKey:IDToko"`
	IDCategory     uint         `gorm:"not null" json:"id_category"`
	Kategori       Kategori     `gorm:"foreignKey:IDCategory"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt      time.Time    `gorm:"autoCreateTime" json:"created_at"`
	FotoProduk     []FotoProduk `gorm:"foreignKey:IDProduk;constraint:OnDelete:CASCADE" json:"foto_produk"`
}
