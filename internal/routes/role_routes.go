package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(rg *gin.RouterGroup, h *handler.RoleHandler) {
	rg.GET("/roles", h.GetAll)
	rg.GET("/roles/:id", h.GetByID)
	rg.POST("/roles", h.Create)
	rg.PATCH("/roles/:id", h.Update)
	rg.DELETE("/roles/:id", h.Delete)
} 