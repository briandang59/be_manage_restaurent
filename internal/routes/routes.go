package routes

import (
	"manage_restaurent/internal/handler"
	"manage_restaurent/internal/repository"
	"manage_restaurent/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"manage_restaurent/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api", middlewares.AuthMiddleware())
	noAuth := r.Group("/auth")

	// Dependencies for Account
	accountRepo := repository.NewAccountRepo(db)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService)
	AccountPublicRoutes(noAuth, accountHandler)
	AccountProtectedRoutes(api, accountHandler)

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

	TableRepo := repository.NewTableRepo(db)
	TableSVC := service.NewTableService(TableRepo)
	TableHandler := handler.NewTableHandler(TableSVC)
	TableRoutes(api, TableHandler)

	fileRepo := repository.NewFileRepo(db)
	fileHandler := handler.NewFileHandler(fileRepo)
	FileRoutes(api, fileHandler)

	// Dependencies for Role
	roleRepo := repository.NewRoleRepo(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)
	RoleRoutes(api, roleHandler)

	// Dependencies for Permission
	permissionRepo := repository.NewPermissionRepo(db)
	permissionService := service.NewPermissionService(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionService)
	PermissionRoutes(api, permissionHandler)

	// Dependencies for Ticket
	ticketRepo := repository.NewTicketRepo(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)
	TicketRoutes(api, ticketHandler)

	// Dependencies for Ingredient
	ingredientRepo := repository.NewIngredientRepo(db)
	ingredientService := service.NewIngredientService(ingredientRepo)
	ingredientHandler := handler.NewIngredientHandler(ingredientService)
	IngredientRoutes(api, ingredientHandler)

	// Dependencies for Attendance
	attendanceRepo := repository.NewAttendanceRepo(db)
	attendanceService := service.NewAttendanceService(attendanceRepo)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	AttendanceRoutes(api, attendanceHandler)

	// Dependencies for OrderItem
	orderItemRepo := repository.NewOrderItemRepo(db)
	orderItemService := service.NewOrderItemService(orderItemRepo)
	orderItemHandler := handler.NewOrderItemHandler(orderItemService)
	OrderItemRoutes(api, orderItemHandler)

	// Dependencies for Order
	orderRepo := repository.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)
	OrderRoutes(api, orderHandler)

	// Dependencies for MenuItem
	menuItemRepo := repository.NewMenuItemRepo(db)
	menuItemService := service.NewMenuItemService(menuItemRepo)
	menuItemHandler := handler.NewMenuItemHandler(menuItemService)
	MenuItemRoutes(api, menuItemHandler)

	// Dependencies for Category
	categoryRepo := repository.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	CategoryRoutes(api, categoryHandler)
}
