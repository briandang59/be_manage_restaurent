package routes

import (
	"manage_restaurent/internal/handler"
	"github.com/gin-gonic/gin"
)

// Đăng ký các route public (không cần token)
func AccountPublicRoutes(rg *gin.RouterGroup, h *handler.AccountHandler) {
	rg.POST("/accounts/login", h.Login)
	rg.POST("/accounts", h.Create)
}

// Đăng ký các route cần xác thực (cần token)
func AccountProtectedRoutes(rg *gin.RouterGroup, h *handler.AccountHandler) {
	rg.GET("/accounts", h.GetAll)
	rg.GET("/accounts/:id", h.GetByID)
	rg.PATCH("/accounts/:id", h.Update)
	rg.DELETE("/accounts/:id", h.Delete)
} 