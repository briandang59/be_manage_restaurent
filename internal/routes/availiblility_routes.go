package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func AvailibilityRoutes(rg *gin.RouterGroup, h *handler.AvailibilityHandler) {
	rg.GET("/availibilities", h.GetAll)
	rg.GET("/availibilities/:id", h.GetByID)
	rg.POST("/availibilities", h.Create)
	rg.PATCH("/availibilities/:id", h.Update)
	rg.DELETE("/availibilities/:id", h.Delete)
}
