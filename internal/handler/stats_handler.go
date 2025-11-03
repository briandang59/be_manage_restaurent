package handler

import (
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	svc *service.StatsService
}

func NewStatsHandler(s *service.StatsService) *StatsHandler {
	return &StatsHandler{svc: s}
}

// RevenueStats godoc
// @Summary Thống kê doanh thu
// @Description Thống kê doanh thu từ đơn hàng
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/revenue [get]
func (h *StatsHandler) RevenueStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.RevenueStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// IngredientsStats godoc
// @Summary Thống kê nguyên liệu
// @Description Thống kê nguyên liệu, bao gồm sắp hết hàng
// @Tags stats
// @Produce json
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/ingredients [get]
func (h *StatsHandler) IngredientsStats(c *gin.Context) {
	stats, err := h.svc.IngredientsStats()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// EmployeesStats godoc
// @Summary Thống kê nhân viên
// @Description Thống kê giờ làm việc của nhân viên
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/employees [get]
func (h *StatsHandler) EmployeesStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.EmployeesStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// OrdersStats godoc
// @Summary Thống kê đơn hàng
// @Description Thống kê đơn hàng và món ăn bán chạy
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/orders [get]
func (h *StatsHandler) OrdersStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.OrdersStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// BookingsStats godoc
// @Summary Thống kê đặt chỗ
// @Description Thống kê đặt chỗ
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/bookings [get]
func (h *StatsHandler) BookingsStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.BookingsStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// CustomersStats godoc
// @Summary Thống kê khách hàng
// @Description Thống kê khách hàng mới
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/customers [get]
func (h *StatsHandler) CustomersStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.CustomersStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}

// TicketsStats godoc
// @Summary Thống kê phiếu nguyên liệu
// @Description Thống kê phiếu nguyên liệu
// @Tags stats
// @Produce json
// @Param from_date query string false "Từ ngày (YYYY-MM-DD)"
// @Param to_date query string false "Đến ngày (YYYY-MM-DD)"
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Router /stats/tickets [get]
func (h *StatsHandler) TicketsStats(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	stats, err := h.svc.TicketsStats(fromDate, toDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, stats, nil)
}
