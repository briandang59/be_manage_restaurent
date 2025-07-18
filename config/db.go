package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Thêm vòng lặp thử lại (retry loop)
	var db *gorm.DB
	var err error
	maxRetries := 10              // Số lần thử lại tối đa
	retryDelay := 5 * time.Second // Thời gian chờ giữa các lần thử lại

	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Kết nối cơ sở dữ liệu thành công")
			DB = db
			return db
		}
		log.Printf("❌ Kết nối cơ sở dữ liệu thất bại (lần %d/%d): %v", i, maxRetries, err)
		time.Sleep(retryDelay)
	}

	// Nếu sau tất cả các lần thử lại mà vẫn thất bại
	log.Fatalf("❌ Kết nối cơ sở dữ liệu thất bại sau %d lần thử. Ứng dụng sẽ thoát.", maxRetries)
	return nil
}
