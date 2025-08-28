package controllers

import (
	"database/sql"
	"net/http"

	"FinalProject/repository"

	"github.com/gin-gonic/gin"
)

type HasilKonversiController struct {
	DB *sql.DB
}

func (ctrl *HasilKonversiController) GetAllLaporanPembelianKonversi(c *gin.Context) {
	result, err := repository.GetAllLaporanPembelianKonversi(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
