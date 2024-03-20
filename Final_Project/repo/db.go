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
	// db.Debug().Migrator().DropTable(&models.User{})
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

func CreateDummyData() {
	db := GetDB()

	// Buat beberapa data dummy untuk pengguna (users)
	users := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1", Age: 25},
		{Username: "user2", Email: "user2@example.com", Password: "password2", Age: 30},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Gagal membuat data pengguna: %v", err)
		}
	}

	// Buat beberapa data dummy untuk foto (photos)
	photos := []models.Photo{
		{Title: "Foto 1", Caption: "Ini adalah foto pertama", PhotoUrl: "http://example.com/photo1.jpg", UserID: 1},
		{Title: "Foto 2", Caption: "Ini adalah foto kedua", PhotoUrl: "http://example.com/photo2.jpg", UserID: 2},
	}

	for _, photo := range photos {
		if err := db.Create(&photo).Error; err != nil {
			log.Fatalf("Gagal membuat data foto: %v", err)
		}
	}

	// Buat beberapa data dummy untuk komentar (comments)
	comments := []models.Comment{
		{Message: "Komentar pertama", PhotoID: 1, UserID: 2},
		{Message: "Komentar kedua", PhotoID: 2, UserID: 1},
	}

	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			log.Fatalf("Gagal membuat data komentar: %v", err)
		}
	}

	// Buat beberapa data dummy untuk media sosial (social media)
	socialMedias := []models.SocialMedia{
		{Name: "Facebook", SocialMediaUrl: "http://facebook.com/user1", UserID: 1},
		{Name: "Instagram", SocialMediaUrl: "http://instagram.com/user2", UserID: 2},
	}

	for _, socialMedia := range socialMedias {
		if err := db.Create(&socialMedia).Error; err != nil {
			log.Fatalf("Gagal membuat data media sosial: %v", err)
		}
	}

	log.Println("Data dummy berhasil dibuat")
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
