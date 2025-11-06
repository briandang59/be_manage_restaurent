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

	// Query cho total_hours tổng (giữ nguyên logic)
	query := s.db.Model(&model.Attendance{})
	if !start.IsZero() {
		query = query.Where("actual_start_time >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("actual_end_time <= ?", end)
	}

	var totalHours int64
	query.Select("SUM(hours) as total").Scan(&totalHours)

	// Query cho employee stats với đầy đủ info và tính salary
	// Giả sử: FULLTIME: lương giờ = (base_salary / 26 / 8) * total_hours
	// PARTTIME: salary_per_hour * total_hours
	// Hardcode 26 ngày/tháng, 8 giờ/ngày cho FULLTIME
	const daysPerMonth = 26
	const hoursPerDay = 8

	type EmployeeStat struct {
		ID            uint                      `json:"id"`
		FullName      string                    `json:"full_name"`
		Gender        bool                      `json:"gender"`
		Birthday      string                    `json:"birthday"`
		PhoneNumber   string                    `json:"phone_number"`
		Email         string                    `json:"email"`
		ScheduleType  enum.EmployeeScheduleType `json:"schedule_type"`
		Address       string                    `json:"address"`
		JoinDate      string                    `json:"join_date"`
		BaseSalary    int64                     `json:"base_salary"`
		SalaryPerHour int64                     `json:"salary_per_hour"`
		AvatarFileID  *uint                     `json:"avatar_file_id"`
		TotalHours    int64                     `json:"total_hours"`
		Salary        int64                     `json:"salary"` // Tính sau khi scan
		// Có thể thêm AvatarFile info nếu JOIN files, nhưng giữ đơn giản
	}
	var employeeStats []EmployeeStat

	empQuery := s.db.Model(&model.Attendance{}).
		Joins("JOIN shift_schedules ON attendances.shift_schedule_id = shift_schedules.id").
		Joins("LEFT JOIN employees ON shift_schedules.employee_id = employees.id") // LEFT nếu cần, nhưng giả sử luôn có
	if !start.IsZero() {
		empQuery = empQuery.Where("attendances.actual_start_time >= ?", start)
	}
	if !end.IsZero() {
		empQuery = empQuery.Where("attendances.actual_end_time <= ?", end)
	}

	// Select tất cả fields non-aggregate từ employees + SUM(hours)
	empQuery.Select(`
		employees.id,
		employees.full_name,
		employees.gender,
		employees.birthday,
		employees.phone_number,
		employees.email,
		employees.schedule_type,
		employees.address,
		employees.join_date,
		employees.base_salary,
		employees.salary_per_hour,
		employees.avatar_file_id,
		SUM(attendances.hours) as total_hours
	`).
		Group(`
		employees.id,
		employees.full_name,
		employees.gender,
		employees.birthday,
		employees.phone_number,
		employees.email,
		employees.schedule_type,
		employees.address,
		employees.join_date,
		employees.base_salary,
		employees.salary_per_hour,
		employees.avatar_file_id
	`).
		Scan(&employeeStats)

	// Tính salary sau khi scan
	for i := range employeeStats {
		stat := &employeeStats[i]
		if stat.TotalHours == 0 {
			stat.Salary = 0
			continue
		}
		switch stat.ScheduleType {
		case "FULLTIME": // Giả sử enum.FullTime tồn tại; thay bằng giá trị thực nếu khác
			if stat.BaseSalary > 0 {
				dailySalary := stat.BaseSalary / daysPerMonth
				hourlySalary := dailySalary / hoursPerDay
				stat.Salary = hourlySalary * stat.TotalHours
			} else {
				stat.Salary = 0
			}
		case "PARTTIME": // Giả sử enum.PartTime tồn tại
			stat.Salary = stat.SalaryPerHour * stat.TotalHours
		default:
			stat.Salary = 0
		}
	}

	return map[string]interface{}{
		"total_hours":    totalHours,
		"employee_stats": employeeStats, // Thay vì employee_hours, dùng stats đầy đủ
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

	// Thêm JOIN với menu_items và files
	query := s.db.Model(&model.OrderItem{}).
		Joins("JOIN menu_items ON order_items.menu_item_id = menu_items.id").
		Joins("LEFT JOIN files ON menu_items.file_id = files.id") // LEFT JOIN vì File có thể optional
	if !start.IsZero() {
		query = query.Where("order_items.created_at >= ?", start)
	}
	if !end.IsZero() {
		query = query.Where("order_items.created_at <= ?", end)
	}

	// Cập nhật struct để chứa thêm info từ File
	type TopItem struct {
		MenuItemID  uint   `json:"menu_item_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
		TotalQty    int64  `json:"total_qty"`
		// Thêm fields từ File (optional, vì LEFT JOIN)
		FileID   *uint  `json:"file_id,omitempty"`
		FileName string `json:"file_name,omitempty"`
		Url      string `json:"url,omitempty"`
		MimeType string `json:"mime_type,omitempty"`
		// Thêm các trường khác từ File nếu cần
	}
	var topItems []TopItem
	query.Select("order_items.menu_item_id, menu_items.name, menu_items.description, menu_items.price, menu_items.file_id, files.file_name, files.url, files.mime_type, SUM(order_items.quantity) as total_qty").
		Group("order_items.menu_item_id, menu_items.name, menu_items.description, menu_items.price, menu_items.file_id, files.file_name, files.url, files.mime_type"). // Group thêm các trường non-aggregate
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
