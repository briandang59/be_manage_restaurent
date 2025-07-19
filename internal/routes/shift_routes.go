package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func ShiftRoutes(rg *gin.RouterGroup, h *handler.ShiftHandler) {
	rg.GET("/shifts", h.GetAll)
	rg.GET("/shifts/:id", h.GetByID)
	rg.POST("/shifts", h.Create)
	rg.PATCH("/shifts/:id", h.Update)
	rg.DELETE("/shifts/:id", h.Delete)
}
