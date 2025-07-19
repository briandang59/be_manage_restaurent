package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func ShiftScheduleRoutes(rg *gin.RouterGroup, h *handler.ShiftScheduleHandler) {
	rg.GET("/shifts-chedules", h.GetAll)
	rg.GET("/shifts-chedules/:id", h.GetByID)
	rg.POST("/shifts-chedules", h.Create)
	rg.PATCH("/shifts-chedules/:id", h.Update)
	rg.DELETE("/shifts-chedules/:id", h.Delete)
}
