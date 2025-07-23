package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"manage_restaurent/config"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/routes"

	docs "manage_restaurent/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title Restaurant Management API
// @version 1.0
// @description API quản lý nhà hàng, nhân viên, thực đơn, đơn hàng, file upload...
// @host localhost:8080
// @BasePath /api
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
		&model.Table{},
		&model.Shift{},
		&model.MenuItem{},
		&model.Account{},
		&model.Role{},
		&model.Permission{},
		&model.Ticket{},
		&model.Ingredient{},
		&model.Attendance{},
		&model.OrderItem{},
		&model.Order{},
		&model.File{},
	); err != nil {
		log.Fatal("❌ AutoMigrate lỗi:", err)
	}

	// Seed roles và permissions chỉ khi lần đầu (nếu chưa có dữ liệu)
	var roleCount int64
	var permCount int64
	db.Model(&model.Role{}).Count(&roleCount)
	db.Model(&model.Permission{}).Count(&permCount)
	if roleCount == 0 || permCount == 0 {
		model.SeedRolesAndPermissions(db)
		log.Println("✅ Đã seed roles và permissions mẫu!")
	}

	// Tạo router
	r := gin.Default()

	// Swagger endpoint
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
