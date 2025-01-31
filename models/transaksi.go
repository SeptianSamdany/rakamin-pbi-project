package models

import (
	"time"
)

type Transaksi struct {
	ID              uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser          uint              `gorm:"not null" json:"id_user"`
	AlamatPengirim  string              `gorm:"not null" json:"alamat_pengirim"`
	HargaTotal      int               `gorm:"not null" json:"harga_total"`
	KodeInvoice     string            `gorm:"unique;not null" json:"kode_invoice"`
	MethodBayar     string            `gorm:"not null" json:"method_bayar"`
	UpdatedAt       time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt       time.Time         `gorm:"autoCreateTime" json:"created_at"`
	DetailTransaksi []DetailTransaksi `gorm:"foreignKey:IDTrx;constraint:OnDelete:CASCADE" json:"detail_transaksi"`
}
