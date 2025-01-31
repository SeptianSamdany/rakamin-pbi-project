package models

import (
	"time"
)

type LogProduk struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk      uint      `gorm:"not null" json:"id_produk"`
	NamaProduk    string    `gorm:"not null" json:"nama_produk"`
	Slug          string    `gorm:"not null" json:"slug"`
	HargaReseller int       `gorm:"not null" json:"harga_reseller"`
	HargaKonsumen int       `gorm:"not null" json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	IDToko        uint      `gorm:"not null" json:"id_toko"`
	IDCategory    uint      `gorm:"not null" json:"id_category"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}