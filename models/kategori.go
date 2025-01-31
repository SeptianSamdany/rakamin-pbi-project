package models

import (
	"time"
)

// Kategori model
type Kategori struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaKategori string    `gorm:"not null" json:"nama_kategori"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
