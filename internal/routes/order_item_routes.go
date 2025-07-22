package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(rg *gin.RouterGroup, h *handler.OrderItemHandler) {
	rg.GET("/order-items", h.GetAll)
	rg.GET("/order-items/:id", h.GetByID)
	rg.POST("/order-items", h.Create)
	rg.PATCH("/order-items/:id", h.Update)
	rg.DELETE("/order-items/:id", h.Delete)
} 