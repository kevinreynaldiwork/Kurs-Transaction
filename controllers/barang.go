package controllers

import (
	"database/sql"
	"net/http"

	"FinalProject/repository"
	"FinalProject/structs"

	"github.com/gin-gonic/gin"
)

// BarangController struct
type BarangController struct {
	DB *sql.DB
}

// GetAllBarang - ambil semua data barang
func (bc *BarangController) GetAllBarang(ctx *gin.Context) {
	result, err := repository.GetAllBarang(bc.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

// GetOneBarang - ambil data barang berdasarkan kode
func (bc *BarangController) GetOneBarang(ctx *gin.Context) {
	kode := ctx.Param("kode")

	result, err := repository.GetOneBarang(bc.DB, kode)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

// CreateBarang - tambah barang baru
func (bc *BarangController) CreateBarang(ctx *gin.Context) {
	var b structs.Barang

	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.InsertBarang(bc.DB, b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusCreated, b)
	ctx.JSON(http.StatusCreated, gin.H{"message": "barang created"})

}

// UpdateBarang - update data barang
func (bc *BarangController) UpdateBarang(ctx *gin.Context) {
	kode := ctx.Param("kode")
	var b structs.Barang

	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b.KodeBarang = kode

	err := repository.UpdateBarang(bc.DB, b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, b)
	ctx.JSON(http.StatusCreated, gin.H{"message": "barang updated"})

}

// DeleteBarang - hapus data barang
func (bc *BarangController) DeleteBarang(ctx *gin.Context) {
	kode := ctx.Param("kode")

	err := repository.DeleteBarang(bc.DB, kode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"deleted": kode})
	ctx.JSON(http.StatusCreated, gin.H{"message": "barang deleted"})

}
