package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"manage_restaurent/config"
	"manage_restaurent/internal/middlewares"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/routes"

	docs "manage_restaurent/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		&model.Category{},
		&model.Booking{},
		&model.Recruitment{},
		&model.ApplyRecruitment{},
	); err != nil {
		log.Fatal("❌ AutoMigrate lỗi:", err)
	}

	// =========================================================
	// 💡 LOGIC CHẠY SEED DATA
	// =========================================================
	var roleCount int64
	var permCount int64
	db.Model(&model.Role{}).Count(&roleCount)
	db.Model(&model.Permission{}).Count(&permCount)

	// Chỉ chạy seeding (Go code và SQL file) nếu database trống
	if roleCount == 0 || permCount == 0 {
		// 1. Chạy seed roles và permissions (Go code)
		model.SeedRolesAndPermissions(db)
		log.Println("✅ Đã seed roles và permissions mẫu!")

		// 2. Chạy file SQL seed data (Dữ liệu lớn từ production/staging)

	} else {
		log.Printf("ℹ️  Database đã có dữ liệu (%d roles, %d perms). Bỏ qua seeding.", roleCount, permCount)
	}
	// =========================================================
	seedFilePath := "go_db_seed_data.sql"
	if err := model.RunSQLSeedFile(db, seedFilePath); err != nil {
		// Lỗi khi thực thi SQL là lỗi nghiêm trọng, dừng ứng dụng
		log.Fatalf("❌ Lỗi nghiêm trọng khi chạy file seed SQL: %v", err)
	}
	// Tạo router
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

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
