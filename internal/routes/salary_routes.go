package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func SalaryRoutes(rg *gin.RouterGroup, h *handler.SalaryHandler) {
	salary := rg.Group("/salary")
	{
		salary.GET("/employees/:employee_id", h.GetSalaryByEmployeeAndMonth)
		salary.GET("/all", h.GetAllSalaries)
	}
}
