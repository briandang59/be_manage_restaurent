package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"manage_restaurent/config"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/routes"
)

func main() {
	// Load biến môi trường từ file .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️  Không tìm thấy file .env, sử dụng biến môi trường hiện tại.")
	}

	// Kết nối DB
	db := config.ConnectDatabase()

	// Auto migrate
	if err := db.AutoMigrate(
		&model.Customer{},
		&model.Employee{},
		&model.Availibility{},
		&model.ShiftSchedule{},
		&model.Shift{}); err != nil {
		log.Fatal("❌ AutoMigrate lỗi:", err)
	}

	// Tạo router
	r := gin.Default()

	// Đăng ký route
	routes.RegisterRoutes(r, db)

	// Lấy port từ biến môi trường
	port := os.Getenv("BE_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server chạy tại http://localhost:%s", port)
	r.Run(":" + port)
}
