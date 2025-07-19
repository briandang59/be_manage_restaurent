package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(rg *gin.RouterGroup, h *handler.CustomerHandler) {
	rg.GET("/customers", h.GetAll)
	rg.POST("/customers", h.Create)
	rg.PUT("/customers/:id", h.Update)
	rg.DELETE("/customers/:id", h.Delete)
}
