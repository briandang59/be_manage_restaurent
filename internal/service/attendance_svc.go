package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"time"
)

type AttendanceService struct {
	repo *repository.AttendanceRepo
}

func NewAttendanceService(r *repository.AttendanceRepo) *AttendanceService {
	return &AttendanceService{repo: r}
}

func (s *AttendanceService) Create(attendance *model.Attendance) error {
	// Tự động tính toán hours từ actual_start_time và actual_end_time
	if !attendance.ActualStartTime.IsZero() && !attendance.ActualEndTime.IsZero() {
		duration := attendance.ActualEndTime.Sub(attendance.ActualStartTime)
		attendance.Hours = int64(duration.Hours())
	}
	return s.repo.Create(attendance)
}

func (s *AttendanceService) GetByID(id uint) (*model.Attendance, error) {
	return s.repo.GetByID(id)
}

func (s *AttendanceService) Update(id uint, updates map[string]interface{}) error {
	var start, end time.Time
	var hasStart, hasEnd bool

	// parse actual_start_time
	if startStr, ok := updates["actual_start_time"].(string); ok {
		parsedStart, err := time.Parse(time.RFC3339, startStr)
		if err == nil {
			start = parsedStart
			updates["actual_start_time"] = parsedStart
			hasStart = true
		}
	}

	// parse actual_end_time
	if endStr, ok := updates["actual_end_time"].(string); ok {
		parsedEnd, err := time.Parse(time.RFC3339, endStr)
		if err == nil {
			end = parsedEnd
			updates["actual_end_time"] = parsedEnd
			hasEnd = true
		}
	}

	// Tự động tính giờ nếu đủ 2 field
	if hasStart && hasEnd && !start.IsZero() && !end.IsZero() {
		duration := end.Sub(start)
		hours := duration.Hours()
		updates["hours"] = hours
	}

	return s.repo.Update(id, updates)
}

func (s *AttendanceService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AttendanceService) List(employeeID *uint, offset, limit int) ([]model.Attendance, int64, error) {
	return s.repo.List(employeeID, offset, limit)
}
