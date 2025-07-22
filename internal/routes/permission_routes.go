package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func PermissionRoutes(rg *gin.RouterGroup, h *handler.PermissionHandler) {
	rg.GET("/permissions", h.GetAll)
	rg.GET("/permissions/:id", h.GetByID)
	rg.POST("/permissions", h.Create)
	rg.PATCH("/permissions/:id", h.Update)
	rg.DELETE("/permissions/:id", h.Delete)
} 