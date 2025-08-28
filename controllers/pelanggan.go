package controllers

import (
	"FinalProject/repository"
	"FinalProject/structs"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PelangganController struct {
	DB *sql.DB
}

// GetAllPelanggan : ambil semua data pelanggan
func (ctrl *PelangganController) GetAllPelanggan(c *gin.Context) {
	result, err := repository.GetAllPelanggan(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetOnePelanggan : ambil data pelanggan berdasarkan kode
func (ctrl *PelangganController) GetOnePelanggan(c *gin.Context) {
	kode := c.Param("kode")

	p, err := repository.GetOnePelanggan(ctrl.DB, kode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// CreatePelanggan : tambah pelanggan baru
func (ctrl *PelangganController) CreatePelanggan(c *gin.Context) {
	var p structs.Pelanggan
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.InsertPelanggan(ctrl.DB, p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "pelanggan created"})
}

// UpdatePelanggan : update data pelanggan
func (ctrl *PelangganController) UpdatePelanggan(c *gin.Context) {
	kode := c.Param("kode")

	var p structs.Pelanggan
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p.KodePelanggan = kode

	if err := repository.UpdatePelanggan(ctrl.DB, p); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pelanggan updated"})
}

// DeletePelanggan : hapus pelanggan
func (ctrl *PelangganController) DeletePelanggan(c *gin.Context) {
	kode := c.Param("kode")

	if err := repository.DeletePelanggan(ctrl.DB, kode); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pelanggan deleted"})
}
