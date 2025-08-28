package routers

import (
	"database/sql"

	"FinalProject/controllers"
	"FinalProject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Controller instances
	matauangCtrl := controllers.MataUangController{DB: db}
	barangCtrl := controllers.BarangController{DB: db}
	kursCtrl := controllers.KursController{DB: db}
	hasilKonversiCtrl := controllers.HasilKonversiController{DB: db}
	pelangganCtrl := &controllers.PelangganController{DB: db}
	transaksiCtrl := controllers.TransaksiController{DB: db}

	// Public routes (Auth)
	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)

	// Protected routes (JWT required)
	api := r.Group("/api", middleware.JWTAuthMiddleware())
	{
		api.GET("/mata-uang", matauangCtrl.GetAllMataUang)
		api.GET("/mata-uang/:kode", matauangCtrl.GetOneMataUang)
		api.POST("/mata-uang", matauangCtrl.CreateMataUang)
		api.PUT("/mata-uang/:kode", matauangCtrl.UpdateMataUang)
		api.DELETE("/mata-uang/:kode", matauangCtrl.DeleteMataUang)

		api.GET("/barang", barangCtrl.GetAllBarang)
		api.GET("/barang/:kode", barangCtrl.GetOneBarang)
		api.POST("/barang", barangCtrl.CreateBarang)
		api.PUT("/barang/:kode", barangCtrl.UpdateBarang)
		api.DELETE("/barang/:kode", barangCtrl.DeleteBarang)

		api.GET("/kurs", kursCtrl.GetAllKurs)
		api.GET("/kurs/:kode/:tanggal", kursCtrl.GetOneKurs)
		api.POST("/kurs", kursCtrl.CreateKurs)
		api.PUT("/kurs/:kode/:tanggal", kursCtrl.UpdateKurs)
		api.DELETE("/kurs/:kode/:tanggal", kursCtrl.DeleteKurs)

		api.GET("/pelanggan", pelangganCtrl.GetAllPelanggan)
		api.GET("/pelanggan/:kode", pelangganCtrl.GetOnePelanggan)
		api.POST("/pelanggan", pelangganCtrl.CreatePelanggan)
		api.PUT("/pelanggan/:kode", pelangganCtrl.UpdatePelanggan)
		api.DELETE("/pelanggan/:kode", pelangganCtrl.DeletePelanggan)

		api.GET("/transaksi", transaksiCtrl.GetAllTransaksi)
		api.GET("/transaksi/:id", transaksiCtrl.GetOneTransaksi)
		api.POST("/transaksi", transaksiCtrl.CreateTransaksi)
		api.PUT("/transaksi/:id", transaksiCtrl.UpdateTransaksi)
		api.DELETE("/transaksi/:id", transaksiCtrl.DeleteTransaksi)

		api.GET("/laporan-konversi", hasilKonversiCtrl.GetAllLaporanPembelianKonversi)

	}

	return r
}
