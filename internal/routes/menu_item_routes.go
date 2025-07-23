package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func MenuItemRoutes(rg *gin.RouterGroup, h *handler.MenuItemHandler) {
	rg.GET("/menu-items", h.GetAll)
	rg.GET("/menu-items/:id", h.GetByID)
	rg.POST("/menu-items", h.Create)
	rg.PATCH("/menu-items/:id", h.Update)
	rg.DELETE("/menu-items/:id", h.Delete)
	rg.POST("/menu-items/import-excel", h.ImportExcel)
} 