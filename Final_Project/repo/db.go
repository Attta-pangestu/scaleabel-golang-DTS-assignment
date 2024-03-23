package repo

import (
	"MyGramAtta/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	port     = "3306"
	dbname   = "my-gram-db"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid connect to database")
	}

	log.Println("Success connect to database")

	db.Debug().AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting underlying database connection: %v", err)
	}
	sqlDB.Close()
}

func CreateDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)

	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal membuat database: %v", err)
	}

	// Buat database jika belum ada
	err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)).Error
	if err != nil {
		log.Fatalf("Gagal membuat database: %v", err)
	}
}

func EnsureTablesCreated() {
	
	// Migrasi model-model untuk membuat tabel-tabel
	db := GetDB()
	err := db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi tabel: %v", err)
	}

	log.Println("Tabel-tabel telah dibuat atau sudah ada")
}
