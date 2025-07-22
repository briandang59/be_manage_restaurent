package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func IngredientRoutes(rg *gin.RouterGroup, h *handler.IngredientHandler) {
	rg.GET("/ingredients", h.GetAll)
	rg.GET("/ingredients/:id", h.GetByID)
	rg.POST("/ingredients", h.Create)
	rg.PATCH("/ingredients/:id", h.Update)
	rg.DELETE("/ingredients/:id", h.Delete)
} 