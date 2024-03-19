package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		var err error
		dsn := "root:@tcp(localhost:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Gagal terhubung ke database")
		}
	}
	return db
}

func Close() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			panic("Gagal mendapatkan koneksi database")
		}
		sqlDB.Close()
	}
}
