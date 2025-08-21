package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(rg *gin.RouterGroup, h *handler.BookingHandler) {
	rg.GET("/bookings", h.GetAll)
	rg.GET("/bookings/:id", h.GetByID)
	rg.POST("/bookings", h.Create)
	rg.PATCH("/bookings/:id", h.Update)
	rg.DELETE("/bookings/:id", h.Delete)
}
