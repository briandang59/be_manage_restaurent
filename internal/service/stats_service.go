package service

import (
	"manage_restaurent/internal/enum"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"time"

	"gorm.io/gorm"
)

type StatsService struct {
	db                *gorm.DB
	orderRepo         *repository.OrderRepo
	ingredientRepo    *repository.IngredientRepo
	attendanceRepo    *repository.AttendanceRepo
	orderItemRepo     *repository.OrderItemRepo
	bookingRepo       *repository.BookingRepo
	customerRepo      *repository.CustomerRepo
	ticketRepo        *repository.TicketRepo
	shiftScheduleRepo *repository.ShiftScheduleRepo
}

func NewStatsService(
	db *gorm.DB,
	orderRepo *repository.OrderRepo,
	ingredientRepo *repository.IngredientRepo,
	attendanceRepo *repository.AttendanceRepo,
	orderItemRepo *repository.OrderItemRepo,
	bookingRepo *repository.BookingRepo,
	customerRepo *repository.CustomerRepo,
	ticketRepo *repository.TicketRepo,
	shiftScheduleRepo *repository.ShiftScheduleRepo,
) *StatsService {
	return &StatsService{
		db:                db,
		orderRepo:         orderRepo,
		ingredientRepo:    ingredientRepo,
		attendanceRepo:    attendanceRepo,
		orderItemRepo:     orderItemRepo,
		bookingRepo:       bookingRepo,
		customerRepo:      customerRepo,
		ticketRepo:        ticketRepo,
		shiftScheduleRepo: shiftScheduleRepo,
	}
}

func parseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", dateStr)
}

func (s *StatsService) RevenueStats(fromDate, toDate string) (map[string]interface{}, error) {
	start, err := parseDate(fromDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDate(toDate)
	if err != nil {
		return nil, err
	}

	query := s.db.Model(&model.Order{}).Where("status = ?", enum.Paid)
	if !start.IsZero() {
		query = query.Where("created_at >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("created_at <= ?", end)
	}

	var totalRevenue int64
	query.Select("SUM(amount) as total").Scan(&totalRevenue)

	var orderCount int64
	query.Count(&orderCount)

	var avgOrder float64
	if orderCount > 0 {
		avgOrder = float64(totalRevenue) / float64(orderCount)
	}

	return map[string]interface{}{
		"total_revenue": totalRevenue,
		"order_count":   orderCount,
		"avg_order":     avgOrder,
	}, nil
}

func (s *StatsService) IngredientsStats() (map[string]interface{}, error) {
	var lowStock []model.Ingredient
	if err := s.db.Where("quantity < warning_quantity").Find(&lowStock).Error; err != nil {
		return nil, err
	}

	var totalIngredients int64
	s.db.Model(&model.Ingredient{}).Count(&totalIngredients)

	type SupplierStat struct {
		Supplier string
		Count    int64
	}
	var supplierStats []SupplierStat
	s.db.Model(&model.Ingredient{}).Select("supplier, COUNT(*) as count").Group("supplier").Scan(&supplierStats)

	return map[string]interface{}{
		"total_ingredients": totalIngredients,
		"low_stock":         lowStock,
		"supplier_stats":    supplierStats,
	}, nil
}

func (s *StatsService) EmployeesStats(fromDate, toDate string) (map[string]interface{}, error) {
	start, err := parseDate(fromDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDate(toDate)
	if err != nil {
		return nil, err
	}

	query := s.db.Model(&model.Attendance{})
	if !start.IsZero() {
		query = query.Where("actual_start_time >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("actual_end_time <= ?", end)
	}

	var totalHours int64
	query.Select("SUM(hours) as total").Scan(&totalHours)

	type EmployeeHour struct {
		EmployeeID uint  `json:"employee_id"`
		TotalHours int64 `json:"total_hours"`
	}
	var employeeHours []EmployeeHour
	s.db.Model(&model.Attendance{}).
		Select("shift_schedules.employee_id, SUM(hours) as total_hours").
		Joins("JOIN shift_schedules ON attendances.shift_schedule_id = shift_schedules.id").
		Group("shift_schedules.employee_id").
		Scan(&employeeHours)

	return map[string]interface{}{
		"total_hours":    totalHours,
		"employee_hours": employeeHours,
	}, nil
}

func (s *StatsService) OrdersStats(fromDate, toDate string) (map[string]interface{}, error) {
	start, err := parseDate(fromDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDate(toDate)
	if err != nil {
		return nil, err
	}

	query := s.db.Model(&model.OrderItem{})
	if !start.IsZero() {
		query = query.Where("created_at >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("created_at <= ?", end)
	}

	type TopItem struct {
		MenuItemID uint  `json:"menu_item_id"`
		TotalQty   int64 `json:"total_qty"`
	}
	var topItems []TopItem
	query.Select("menu_item_id, SUM(quantity) as total_qty").
		Group("menu_item_id").
		Order("total_qty DESC").
		Limit(10).
		Scan(&topItems)

	var totalOrders int64
	s.db.Model(&model.Order{}).Count(&totalOrders)

	return map[string]interface{}{
		"total_orders": totalOrders,
		"top_items":    topItems,
	}, nil
}

func (s *StatsService) BookingsStats(fromDate, toDate string) (map[string]interface{}, error) {
	query := s.db.Model(&model.Booking{})
	if fromDate != "" {
		query = query.Where("booking_date >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("booking_date <= ?", toDate)
	}

	var totalBookings int64
	query.Count(&totalBookings)

	var totalPersons int
	query.Select("SUM(total_persons) as total").Scan(&totalPersons)

	type StatusStat struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}
	var statusStats []StatusStat
	query.Select("status, COUNT(*) as count").Group("status").Scan(&statusStats)

	return map[string]interface{}{
		"total_bookings": totalBookings,
		"total_persons":  totalPersons,
		"status_stats":   statusStats,
	}, nil
}

func (s *StatsService) CustomersStats(fromDate, toDate string) (map[string]interface{}, error) {
	start, err := parseDate(fromDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDate(toDate)
	if err != nil {
		return nil, err
	}

	query := s.db.Model(&model.Customer{})
	if !start.IsZero() {
		query = query.Where("created_at >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("created_at <= ?", end)
	}

	var totalCustomers int64
	query.Count(&totalCustomers)

	return map[string]interface{}{
		"total_customers": totalCustomers,
	}, nil
}

func (s *StatsService) TicketsStats(fromDate, toDate string) (map[string]interface{}, error) {
	start, err := parseDate(fromDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDate(toDate)
	if err != nil {
		return nil, err
	}

	query := s.db.Model(&model.Ticket{})
	if !start.IsZero() {
		query = query.Where("created_at >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("created_at <= ?", end)
	}

	type TypeStat struct {
		TicketType enum.TicketType `json:"ticket_type"`
		TotalQty   int64           `json:"total_qty"`
	}
	var typeStats []TypeStat
	query.Select("ticket_type, SUM(quantity) as total_qty").
		Group("ticket_type").
		Scan(&typeStats)

	var totalTickets int64
	query.Count(&totalTickets)

	return map[string]interface{}{
		"total_tickets": totalTickets,
		"type_stats":    typeStats,
	}, nil
}
