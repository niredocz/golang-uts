package models

type Perlengkapan_Bayi struct {
	Id_Penjualan string `json:"id_penjualan" gorm:"primary_key"`
	Nama_Barang string `json:"nama_barang"`
	Jumlah_Stok int `json:"jumlah_stok"`
	Harga int `json:"harga"`
}

