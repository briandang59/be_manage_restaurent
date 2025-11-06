package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func ApplyRecruitmentRoutes(rg *gin.RouterGroup, h *handler.ApplyRecruitmentHandler) {
	rg.GET("/apply-recruitments", h.GetAll)
	rg.POST("/apply-recruitments", h.Create)
	rg.PATCH("/apply-recruitments/:id", h.Update)
	rg.DELETE("/apply-recruitments/:id", h.Delete)
}
