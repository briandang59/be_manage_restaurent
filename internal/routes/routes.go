package routes

import (
	"manage_restaurent/internal/handler"
	"manage_restaurent/internal/repository"
	"manage_restaurent/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")

	// Dependencies for Customer
	customerRepo := repository.NewCustomerRepo(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)
	CustomerRoutes(api, customerHandler)

	// Dependencies for Shift
	shiftRepo := repository.NewShiftRepo(db)
	shiftService := service.NewShiftService(shiftRepo)
	shiftHandler := handler.NewShiftHandler(shiftService)
	ShiftRoutes(api, shiftHandler)

	// Dependencies for Employee
	employeeRepo := repository.NewEmployeeRepo(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	EmployeeRoutes(api, employeeHandler)

	// Dependencies for Availibility
	availibilityRepo := repository.NewAvailibilityRepo(db)
	availibilityService := service.NewAvailibilityService(availibilityRepo)
	availibilityHandler := handler.NewAvailibilityHandler(availibilityService)
	AvailibilityRoutes(api, availibilityHandler)

	// Dependencies for ShiftSchedule
	shiftScheduleRepo := repository.NewShiftScheduleRepo(db)
	shiftScheduleService := service.NewShiftScheduleService(shiftScheduleRepo)
	shiftScheduleHandler := handler.NewShiftScheduleHandler(shiftScheduleService)
	ShiftScheduleRoutes(api, shiftScheduleHandler)
}
