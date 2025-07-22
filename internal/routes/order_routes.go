package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(rg *gin.RouterGroup, h *handler.OrderHandler) {
	rg.GET("/orders", h.GetAll)
	rg.GET("/orders/:id", h.GetByID)
	rg.POST("/orders", h.Create)
	rg.PATCH("/orders/:id", h.Update)
	rg.DELETE("/orders/:id", h.Delete)
} 