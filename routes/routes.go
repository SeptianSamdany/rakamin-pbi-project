package routes

import (
	"rakamin-project/controllers"
	"rakamin-project/middlewares"
	"rakamin-project/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Inisialisasi services
	wilayahService := services.NewWilayahService()
	userService := services.NewUserService(db, wilayahService)
	tokoService := services.NewTokoService(db)
	kategoriService := services.NewKategoriService(db)
	produkService := services.NewProdukService(db)
	alamatService := services.NewAlamatService(db)
	transaksiService := services.NewTransaksiService(db)

	// Inisialisasi controllers
	userController := controllers.UserController{UserService: userService}
	tokoController := controllers.TokoController{TokoService: tokoService}
	kategoriController := controllers.KategoriController{KategoriService: kategoriService}
	produkController := controllers.ProdukController{ProdukService: produkService}
	alamatController := controllers.AlamatController{AlamatService: alamatService}
	transaksiController := controllers.TransaksiController{TransaksiService: transaksiService}

	// Rute untuk user
	userRoutes := app.Group("/users")
	userRoutes.Post("/register", userController.Register)
	userRoutes.Post("/login", userController.Login)
	userRoutes.Put("/", middlewares.AuthMiddleware, userController.UpdateUser)
	userRoutes.Delete("/", middlewares.AuthMiddleware, userController.DeleteUser)

	// Rute untuk toko
	tokoRoutes := app.Group("/toko")
	tokoRoutes.Get("/", middlewares.AuthMiddleware, tokoController.GetToko)
	tokoRoutes.Put("/", middlewares.AuthMiddleware, tokoController.UpdateToko)

		// Rute untuk kategori (Admin Only)
	kategoriRoutes := app.Group("/kategori")

	// AuthMiddleware harus dijalankan dulu sebelum AdminMiddleware
	kategoriRoutes.Use(middlewares.AuthMiddleware)
	kategoriRoutes.Use(middlewares.AdminMiddleware)

	kategoriRoutes.Post("/", kategoriController.CreateKategori)
	kategoriRoutes.Get("/", kategoriController.GetAllKategori)
	kategoriRoutes.Put("/:id", kategoriController.UpdateKategori)
	kategoriRoutes.Delete("/:id", kategoriController.DeleteKategori)


	// Rute untuk produk
	produkRoutes := app.Group("/produk")
	produkRoutes.Post("/", middlewares.AuthMiddleware, produkController.CreateProduk)
	produkRoutes.Get("/", produkController.GetAllProduk)
	produkRoutes.Get("/:id", produkController.GetProdukByID)
	produkRoutes.Put("/:id", middlewares.AuthMiddleware, produkController.UpdateProduk)
	produkRoutes.Delete("/:id", middlewares.AuthMiddleware, produkController.DeleteProduk)

	// Rute untuk alamat
	alamatRoutes := app.Group("/alamat")
	alamatRoutes.Post("/", middlewares.AuthMiddleware, alamatController.CreateAlamat)
	alamatRoutes.Get("/", middlewares.AuthMiddleware, alamatController.GetAllAlamat)
	alamatRoutes.Get("/:id", middlewares.AuthMiddleware, alamatController.GetAlamatByID)
	alamatRoutes.Put("/:id", middlewares.AuthMiddleware, alamatController.UpdateAlamat)
	alamatRoutes.Delete("/:id", middlewares.AuthMiddleware, alamatController.DeleteAlamat)

	// Rute untuk transaksi
	transaksiRoutes := app.Group("/transaksi")
	transaksiRoutes.Post("/", middlewares.AuthMiddleware, transaksiController.CreateTransaksi)
	transaksiRoutes.Get("/:id", middlewares.AuthMiddleware, transaksiController.GetTransaksiByID)
	transaksiRoutes.Put("/:id", middlewares.AuthMiddleware, transaksiController.UpdateTransaksi)
	transaksiRoutes.Delete("/:id", middlewares.AuthMiddleware, transaksiController.DeleteTransaksi)
	transaksiRoutes.Post("/detail", middlewares.AuthMiddleware, transaksiController.SaveDetailTransaksi)
	transaksiRoutes.Post("/log-produk", middlewares.AuthMiddleware, transaksiController.CreateLogProduk)
}
