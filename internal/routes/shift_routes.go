package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func ShiftRoutes(rg *gin.RouterGroup, h *handler.ShiftHandler) {
	rg.GET("/shift", h.GetAll)
	rg.GET("/shift/:id", h.GetByID)
	rg.POST("/shift", h.Create)
	rg.PATCH("/shift/:id", h.Update)
	rg.DELETE("/shift/:id", h.Delete)
}
