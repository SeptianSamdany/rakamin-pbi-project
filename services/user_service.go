package services

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"rakamin-project/models"
	"fmt"
)

type UserService struct {
	DB *gorm.DB
	WilayahService *WilayahService
}

func NewUserService(db *gorm.DB, wilayahService *WilayahService) *UserService {
	return &UserService{DB: db, WilayahService: wilayahService}
}

// Register user baru
func (us *UserService) Register(userInput models.User) (*models.User, error) {
	var existingUser models.User
	if err := us.DB.Where("email = ? OR no_telp = ?", userInput.Email, userInput.NoTelp).First(&existingUser).Error; err == nil {
		return nil, errors.New("email atau no telepon sudah terdaftar")
	}

	// Validasi ID Provinsi
	provinsiList, err := us.WilayahService.GetProvinsi()
	if err != nil {
		return nil, errors.New("gagal mengambil data provinsi")
	}

	validProvinsi := false
	for _, prov := range provinsiList {
		if prov.ID == userInput.IDProvinsi {
			validProvinsi = true
			break
		}
	}
	if !validProvinsi {
		return nil, errors.New("ID Provinsi tidak valid")
	}

	// Validasi ID Kota
	kotaList, err := us.WilayahService.GetKotaByProvinsi(userInput.IDProvinsi)
	if err != nil {
		return nil, errors.New("gagal mengambil data kota")
	}

	validKota := false
	for _, kota := range kotaList {
		if kota.ID == userInput.IDKota {
			validKota = true
			break
		}
	}
	if !validKota {
		return nil, errors.New("ID Kota tidak valid")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userInput.KataSandi = string(hashedPassword)

	// Mulai transaksi
	tx := us.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Simpan user baru
	if err := tx.Create(&userInput).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Buat nama toko otomatis
	urlToko := fmt.Sprintf("toko-%s", strings.ToLower(strings.ReplaceAll(userInput.Nama, " ", "-")))
	toko := models.Toko{
		IDUser:   userInput.ID,
		NamaToko: userInput.Nama,
		UrlToko:  urlToko,
	}

	if err := tx.Create(&toko).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	// Ambil data user yang sudah termasuk toko
	us.DB.Preload("Toko").First(&userInput, userInput.ID)
	userInput.KataSandi = "" // Jangan kirim password dalam response

	return &userInput, nil
}

// Login user
func (us *UserService) Login(email, password string) (string, error) {
	var user models.User

	if err := us.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(password)); err != nil {
		return "", errors.New("email atau password salah")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Update user
func (us *UserService) UpdateUser(userID uint, userInput models.User) (*models.User, error) {
	var existingUser models.User

	// Cek apakah user ada
	if err := us.DB.First(&existingUser, userID).Error; err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Jika password diupdate, hash ulang password
	if userInput.KataSandi != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		userInput.KataSandi = string(hashedPassword)
	}

	// Update data user
	if err := us.DB.Model(&existingUser).Updates(userInput).Error; err != nil {
		return nil, err
	}

	updatedUser := existingUser
	updatedUser.KataSandi = "" // Jangan kirimkan kata sandi

	return &updatedUser, nil
}

// Delete user
func (us *UserService) DeleteUser(userID uint) error {
	var user models.User

	// Cari user berdasarkan ID
	if err := us.DB.First(&user, userID).Error; err != nil {
		return errors.New("user tidak ditemukan")
	}

	// Mulai transaksi untuk menghapus user
	tx := us.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Hapus data terkait seperti produk, toko, dll
	if err := tx.Where("id_user = ?", userID).Delete(&models.Toko{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// Fungsi untuk membuat JWT token
func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"is_admin": user.IsAdmin,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "secret_jwt_key" // Ganti dengan key yang lebih aman
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
