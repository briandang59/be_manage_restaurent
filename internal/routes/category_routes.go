package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rg *gin.RouterGroup, h *handler.CategoryHandler) {
	rg.GET("/categories", h.GetAll)
	rg.GET("/categories/:id", h.GetByID)
	rg.POST("/categories", h.Create)
	rg.PATCH("/categories/:id", h.Update)
	rg.DELETE("/categories/:id", h.Delete)
}
