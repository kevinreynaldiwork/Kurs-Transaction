package controllers

import (
	"FinalProject/repository"
	"FinalProject/structs"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransaksiController struct {
	DB *sql.DB
}

// GetAllTransaksi : ambil semua data transaksi
func (ctrl *TransaksiController) GetAllTransaksi(c *gin.Context) {
	result, err := repository.GetAllTransaksi(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetOneTransaksi : ambil data transaksi berdasarkan ID
func (ctrl *TransaksiController) GetOneTransaksi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	t, err := repository.GetOneTransaksi(ctrl.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

// CreateTransaksi : tambah transaksi baru
func (ctrl *TransaksiController) CreateTransaksi(c *gin.Context) {
	var t structs.Transaksi
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.InsertTransaksi(ctrl.DB, t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "transaksi created"})
}

// UpdateTransaksi : update data transaksi
func (ctrl *TransaksiController) UpdateTransaksi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var t structs.Transaksi
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.ID = id

	if err := repository.UpdateTransaksi(ctrl.DB, t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaksi updated"})
}

// DeleteTransaksi : hapus transaksi
func (ctrl *TransaksiController) DeleteTransaksi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := repository.DeleteTransaksi(ctrl.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaksi deleted"})
}
