package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(rg *gin.RouterGroup, h *handler.CustomerHandler) {
	rg.GET("/customer", h.GetAll)
	rg.POST("/customer", h.Create)
	rg.PUT("/customer/:id", h.Update)
	rg.DELETE("/customer/:id", h.Delete)
}
