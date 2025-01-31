package models

import (
	"time"
)

// User model
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string    `gorm:"not null" json:"nama"`
	KataSandi    string    `gorm:"not null" json:"kata_sandi"`
	NoTelp       string    `gorm:"unique;not null" json:"no_telp"`
	TanggalLahir time.Time `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"not null" json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `gorm:"unique;not null" json:"email"`
	IDKota       string    `gorm:"not null" json:"id_kota"`
	IDProvinsi   string    `gorm:"not null" json:"id_provinsi"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	Toko         Toko      `gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE"`
	Alamat       []Alamat 		`gorm:"foreignKey:IDUser;constraint:OnDelete:CASCADE" json:"alamat"`
}
