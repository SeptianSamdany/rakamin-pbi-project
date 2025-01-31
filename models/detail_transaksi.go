package models

import (
	"time"
)

type DetailTransaksi struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDTrx       uint      `gorm:"not null" json:"id_trx"`
	IDLogProduk uint      `gorm:"not null" json:"id_log_produk"`
	IDToko      uint      `gorm:"not null" json:"id_toko"`
	Kuantitas   int       `gorm:"not null" json:"kuantitas"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
