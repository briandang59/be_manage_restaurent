package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(rg *gin.RouterGroup, h *handler.EmployeeHandler) {
	rg.GET("/employees", h.GetAll)
	rg.GET("/employees/:id", h.GetByID)
	rg.POST("/employees", h.Create)
	rg.POST("/employees/with-account", h.CreateWithAutoAccount)
	rg.PATCH("/employees/:id", h.Update)
	rg.DELETE("/employees/:id", h.Delete)
}
