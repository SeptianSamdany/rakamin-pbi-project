# Golang Mini Project

## Deskripsi

Proyek ini adalah API e-commerce sederhana yang dibuat menggunakan **Golang**, **Fiber**, dan **GORM** dengan database **MySQL**. API ini menyediakan fitur untuk mengelola pengguna, toko, kategori, produk, transaksi, dan alamat.

## Teknologi yang Digunakan

- **Golang** - Bahasa pemrograman utama
- **Fiber** - Web framework untuk Golang
- **GORM** - ORM untuk mengelola database MySQL
- **MySQL** - Database yang digunakan
- **Postman** - Untuk testing API
- **JWT (JSON Web Token)** - Untuk autentikasi dan otorisasi

## Instalasi & Menjalankan Aplikasi

### 1. Clone Repository

```bash
git clone https://github.com/SeptianSamdany/golang-mini-project.git
cd golang-mini-project
```

### 2. Buat dan Konfigurasi `.env`

Buat file `.env` dan sesuaikan dengan konfigurasi database kamu:

```
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=password
DB_NAME=rakamin_project
DB_PORT=3306
JWT_SECRET=your_secret_key
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Jalankan Aplikasi

```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080`

## Struktur Folder

```
├── controllers   # Handler untuk setiap endpoint
├── middlewares   # Middleware untuk autentikasi dan otorisasi
├── models        # Struktur tabel database
├── routes        # Routing API
├── services      # Logika bisnis untuk setiap fitur
├── utils         # Utility functions seperti hash dan JWT
├── main.go       # Entry point aplikasi
└── .env          # Konfigurasi database dan JWT
```

## Fitur API

### 1. **Autentikasi (Users)**

- **Register**: `POST /users/register`
- **Login**: `POST /users/login`
- **Update User**: `PUT /users`
- **Delete User**: `DELETE /users`

### 2. **Toko**

- **Get Toko**: `GET /toko`
- **Update Toko**: `PUT /toko`

### 3. **Kategori (Admin Only)**

- **Create Kategori**: `POST /kategori`
- **Get All Kategori**: `GET /kategori`
- **Update Kategori**: `PUT /kategori/:id`
- **Delete Kategori**: `DELETE /kategori/:id`

### 4. **Produk**

- **Create Produk**: `POST /produk`
- **Get All Produk**: `GET /produk`
- **Get Produk by ID**: `GET /produk/:id`
- **Update Produk**: `PUT /produk/:id`
- **Delete Produk**: `DELETE /produk/:id`

### 5. **Foto Produk**

- **Upload Foto Produk**: `POST /produk/:id/upload-foto`
- **Update Foto Produk**: `PUT /produk/:id/update-foto`
- **Delete Foto Produk**: `DELETE /produk/:id/delete-foto`

### 6. **Alamat Pengiriman**

- **Create Alamat**: `POST /alamat`
- **Get All Alamat**: `GET /alamat`
- **Get Alamat by ID**: `GET /alamat/:id`
- **Update Alamat**: `PUT /alamat/:id`
- **Delete Alamat**: `DELETE /alamat/:id`

### 7. **Transaksi**

- **Create Transaksi**: `POST /transaksi`
- **Get Transaksi by ID**: `GET /transaksi/:id`
- **Update Transaksi**: `PUT /transaksi/:id`
- **Delete Transaksi**: `DELETE /transaksi/:id`
- **Save Detail Transaksi**: `POST /transaksi/detail`
- **Create Log Produk**: `POST /transaksi/log-produk`
- **Update Log Produk**: `PUT /transaksi/log-produk/:id`
- **Delete Log Produk**: `DELETE /transaksi/log-produk/:id`

## Pengujian API dengan Postman

1. **Buka Postman** dan impor collection API (jika ada)
2. **Jalankan request API** berdasarkan endpoint di atas
3. **Pastikan response sesuai dengan yang diharapkan**

## Lisensi

Proyek ini dibuat untuk keperluan pembelajaran dan bebas digunakan oleh siapa saja.

---

Jika ada pertanyaan atau kontribusi, silakan ajukan issue atau pull request di repository ini! 🚀

