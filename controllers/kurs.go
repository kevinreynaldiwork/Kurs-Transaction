package controllers

import (
	"FinalProject/repository"
	"FinalProject/structs"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KursController struct {
	DB *sql.DB
}

// GetAllKurs : ambil semua data kurs
func (ctrl *KursController) GetAllKurs(c *gin.Context) {
	result, err := repository.GetAllKurs(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetOneKurs : ambil kurs berdasarkan kode_mata_uang & tanggal
func (ctrl *KursController) GetOneKurs(c *gin.Context) {
	kode := c.Param("kode")
	tanggal := c.Param("tanggal")

	k, err := repository.GetOneKurs(ctrl.DB, kode, tanggal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, k)
}

// CreateKurs : tambah kurs baru
func (ctrl *KursController) CreateKurs(c *gin.Context) {
	var k structs.Kurs
	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.InsertKurs(ctrl.DB, k); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "kurs created"})
}

// UpdateKurs : update kurs
func (ctrl *KursController) UpdateKurs(c *gin.Context) {
	kode := c.Param("kode")
	tanggal := c.Param("tanggal")

	var k structs.Kurs
	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	k.KodeMataUang = kode
	k.Tanggal = tanggal

	if err := repository.UpdateKurs(ctrl.DB, k); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "kurs updated"})
}

// DeleteKurs : hapus kurs
func (ctrl *KursController) DeleteKurs(c *gin.Context) {
	kode := c.Param("kode")
	tanggal := c.Param("tanggal")

	if err := repository.DeleteKurs(ctrl.DB, kode, tanggal); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "kurs deleted"})
}
