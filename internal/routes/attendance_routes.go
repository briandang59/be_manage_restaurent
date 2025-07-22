package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(rg *gin.RouterGroup, h *handler.AttendanceHandler) {
	rg.GET("/attendances", h.GetAll)
	rg.GET("/attendances/:id", h.GetByID)
	rg.POST("/attendances", h.Create)
	rg.PATCH("/attendances/:id", h.Update)
	rg.DELETE("/attendances/:id", h.Delete)
} 