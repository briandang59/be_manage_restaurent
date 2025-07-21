package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func TableRoutes(rg *gin.RouterGroup, h *handler.TableHanlder) {
	rg.GET("/tables", h.GetAll)
	rg.POST("/tables", h.Create)
	rg.PUT("/tables/:id", h.Update)
	rg.DELETE("/tables/:id", h.Delete)
}
