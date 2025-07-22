package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(rg *gin.RouterGroup, h *handler.TicketHandler) {
	rg.GET("/tickets", h.GetAll)
	rg.GET("/tickets/:id", h.GetByID)
	rg.POST("/tickets", h.Create)
	rg.PATCH("/tickets/:id", h.Update)
	rg.DELETE("/tickets/:id", h.Delete)
} 