package controllers

import (
	"database/sql"
	"net/http"

	"FinalProject/repository"
	"FinalProject/structs"

	"github.com/gin-gonic/gin"
)

// MataUangController struct
type MataUangController struct {
	DB *sql.DB
}

func (mc *MataUangController) GetAllMataUang(ctx *gin.Context) {
	result, err := repository.GetAllMataUang(mc.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

func (mc *MataUangController) GetOneMataUang(ctx *gin.Context) {
	kode := ctx.Param("kode")

	result, err := repository.GetOneMataUang(mc.DB, kode)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Mata uang tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

func (mc *MataUangController) CreateMataUang(ctx *gin.Context) {
	var mu structs.MataUang

	// Validasi input
	if err := ctx.ShouldBindJSON(&mu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan lewat repository
	err := repository.InsertMataUang(mc.DB, mu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusCreated, mu)
	ctx.JSON(http.StatusCreated, gin.H{"message": "mata uang created"})

}

func (mc *MataUangController) UpdateMataUang(ctx *gin.Context) {
	var mu structs.MataUang
	kode := ctx.Param("kode")

	if err := ctx.ShouldBindJSON(&mu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.KodeMataUang = kode

	err := repository.UpdateMataUang(mc.DB, mu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, mu)
	ctx.JSON(http.StatusCreated, gin.H{"message": "mata uang updated"})

}

func (mc *MataUangController) DeleteMataUang(ctx *gin.Context) {
	kode := ctx.Param("kode")

	err := repository.DeleteMataUang(mc.DB, kode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"deleted": kode})
	ctx.JSON(http.StatusCreated, gin.H{"message": "mata uang deleted"})

}
