package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-uts/models"
	"golang-uts/controllers"
)

func main() {
	ginDefault := gin.Default();

	//Import Models
	db := models.SetupModels();

    ginDefault.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Routes Start //
	ginDefault.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data":"Sistem Penjualan Perlengkapan Bayi API"})
	})
	
	ginDefault.GET("/penjualan", controllers.TampilPenjualan)
	ginDefault.POST("/penjualan", controllers.TambahPenjualan)
	// Routes End //

	ginDefault.Run()
}

