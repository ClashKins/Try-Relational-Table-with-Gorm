package models

type Mart struct {
	Bulan   string `gorm:"varchar(100)"`
	Makanan string `gorm:"varchar(100)"`
	Minuman string `gorm:"varchar(100)"`
	Harga   int    `gorm:"integer(11)"`
	Id      uint   `gorm:"primaryKey"`
}

type Rent struct {
	Lama_Sewa string `gorm:"varchar(100)"`
	Harga     int    `gorm:"integer(11)"`
	Ukuran    string `gorm:"varchar(100)"`
}