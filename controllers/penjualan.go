package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"golang-uts/models"
)

type PenjualanInput struct {
	Id_Penjualan string `json:"id_penjualan" gorm:"primary_key"`
	Nama_Barang string `json:"nama_barang"`
	Jumlah_Stok int `json:"jumlah_stok"`
	Harga int `json:"harga"`
}

func TampilPenjualan (c *gin.Context) {
	db:= c.MustGet("db").(*gorm.DB)

	var penjualan []models.Perlengkapan_Bayi
	db.Find(&penjualan)
	c.JSON(http.StatusOK, gin.H{"data" : penjualan})
}

func TambahPenjualan (c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Input Validation
	var dataInput PenjualanInput
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Input Process
	penjualan := models.Perlengkapan_Bayi {
		Id_Penjualan: dataInput.Id_Penjualan,
		Nama_Barang: dataInput.Nama_Barang,
		Jumlah_Stok: dataInput.Jumlah_Stok,
		Harga: dataInput.Harga,
	}

	db.Create(&penjualan)

	c.JSON(http.StatusOK, gin.H{"data":penjualan})
}
