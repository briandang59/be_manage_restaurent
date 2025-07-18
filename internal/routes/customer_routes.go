package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func DepartmentRoutes(rg *gin.RouterGroup, h *handler.CustomerHandler) {
	rg.GET("/customer", h.GetAll)

}
