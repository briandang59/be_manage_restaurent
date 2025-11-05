package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func ApplyRecruitmentRoutes(rg *gin.RouterGroup, h *handler.ApplyRecruitmentHandler) {
	apply := rg.Group("/apply-recruitments")
	{
		apply.GET("/", h.GetAll)
		apply.POST("/", h.Create)
		apply.PATCH("/:id", h.Update)
		apply.DELETE("/:id", h.Delete)
	}
}
