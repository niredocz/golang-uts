package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/golanguts?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Gagal Koneksi ke Database! Mohon Periksa Kembali")
	}

	db.AutoMigrate(&Perlengkapan_Bayi{})

	return db
}


