package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func StatsRoutes(rg *gin.RouterGroup, h *handler.StatsHandler) {
	stats := rg.Group("/stats")
	{
		stats.GET("/revenue", h.RevenueStats)
		stats.GET("/ingredients", h.IngredientsStats)
		stats.GET("/employees", h.EmployeesStats)
		stats.GET("/orders", h.OrdersStats)
		stats.GET("/bookings", h.BookingsStats)
		stats.GET("/customers", h.CustomersStats)
		stats.GET("/tickets", h.TicketsStats)
	}
}
