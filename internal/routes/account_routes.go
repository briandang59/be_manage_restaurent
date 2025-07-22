package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(rg *gin.RouterGroup, h *handler.AccountHandler) {
	rg.GET("/accounts", h.GetAll)
	rg.GET("/accounts/:id", h.GetByID)
	rg.POST("/accounts", h.Create)
	rg.PATCH("/accounts/:id", h.Update)
	rg.DELETE("/accounts/:id", h.Delete)
	rg.POST("/accounts/login", h.Login)
} 