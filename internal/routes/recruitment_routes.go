package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func RecruitmentRoutes(rg *gin.RouterGroup, h *handler.RecruitmentHandler) {
	rg.GET("/recruitments", h.GetAll)
	rg.POST("/recruitments", h.Create)
	rg.PATCH("/recruitments/:id", h.Update)
	rg.DELETE("/recruitments/:id", h.Delete)
}
