package models

import (
	"time"
)

// Toko model
type Toko struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser     uint      `gorm:"not null" json:"id_user"`
	NamaToko   string    `gorm:"not null" json:"nama_toko"`
	UrlToko    string    `gorm:"not null;unique" json:"url_toko"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
