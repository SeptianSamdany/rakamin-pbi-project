package models

import "time"

// Alamat model
type Alamat struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser       uint      `gorm:"not null" json:"id_user"`
	JudulAlamat  string    `gorm:"not null" json:"judul_alamat"`   // Contoh: "Rumah", "Kantor"
	NamaPenerima string    `gorm:"not null" json:"nama_penerima"`  // Nama penerima paket
	NoTelp       string    `gorm:"not null" json:"no_telp"`
	DetailAlamat string    `gorm:"not null" json:"detail_alamat"`  // Detail alamat lengkap
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	User         User      `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE" json:"user"` // Relasi ke User
}
