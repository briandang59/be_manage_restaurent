package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"time"

	"gorm.io/gorm"
)

type SalaryService struct {
	db                *gorm.DB
	attendanceRepo    *repository.AttendanceRepo
	shiftScheduleRepo *repository.ShiftScheduleRepo
	// employeeRepo      *repository.EmployeeRepo // Nếu cần lấy rate từ Employee
}

func NewSalaryService(
	db *gorm.DB,
	attendanceRepo *repository.AttendanceRepo,
	shiftScheduleRepo *repository.ShiftScheduleRepo,
	// employeeRepo *repository.EmployeeRepo,
) *SalaryService {
	return &SalaryService{
		db:                db,
		attendanceRepo:    attendanceRepo,
		shiftScheduleRepo: shiftScheduleRepo,
		// employeeRepo:      employeeRepo,
	}
}
func (s *SalaryService) CalculateSalary(employeeID uint, month string) (map[string]interface{}, error) {
	// Parse month to start and end date
	startTime, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, err
	}
	endTime := startTime.AddDate(0, 1, 0) // Next month
	var totalHours int64
	query := s.db.Model(&model.Attendance{}).
		Joins("JOIN shift_schedules ON attendances.shift_schedule_id = shift_schedules.id").
		Where("shift_schedules.employee_id = ?", employeeID).
		Where("actual_start_time >= ? AND actual_start_time < ?", startTime, endTime).
		Select("SUM(hours) as total_hours").
		Scan(&totalHours)
	if query.Error != nil {
		return nil, query.Error
	}
	// Giả sử hourly rate = 100000 VND (có thể lấy từ Employee nếu có)
	hourlyRate := int64(100000) // Thay bằng logic lấy từ DB nếu cần
	salary := totalHours * hourlyRate
	return map[string]interface{}{
		"employee_id":  employeeID,
		"month":        month,
		"total_hours":  totalHours,
		"hourly_rate":  hourlyRate,
		"total_salary": salary,
	}, nil
}

func (s *SalaryService) CalculateAllSalaries(month string) ([]map[string]interface{}, error) {
	startTime, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, err
	}
	endTime := startTime.AddDate(0, 1, 0)

	type Result struct {
		EmployeeID uint
		TotalHours int64
	}
	var results []Result

	// Group theo employee_id để tính tổng giờ
	query := s.db.Table("attendances").
		Joins("JOIN shift_schedules ON attendances.shift_schedule_id = shift_schedules.id").
		Where("actual_start_time >= ? AND actual_start_time < ?", startTime, endTime).
		Select("shift_schedules.employee_id as employee_id, SUM(hours) as total_hours").
		Group("shift_schedules.employee_id").
		Scan(&results)
	if query.Error != nil {
		return nil, query.Error
	}

	hourlyRate := int64(100000) // Giữ nguyên hoặc lấy từ DB
	salaries := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		salaries = append(salaries, map[string]interface{}{
			"employee_id":  r.EmployeeID,
			"month":        month,
			"total_hours":  r.TotalHours,
			"hourly_rate":  hourlyRate,
			"total_salary": r.TotalHours * hourlyRate,
		})
	}
	return salaries, nil
}
