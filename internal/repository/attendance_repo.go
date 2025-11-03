package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type AttendanceRepo struct {
	db *gorm.DB
}

func NewAttendanceRepo(db *gorm.DB) *AttendanceRepo {
	return &AttendanceRepo{db: db}
}

func (r *AttendanceRepo) Create(attendance *model.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepo) GetByID(id uint) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := r.db.Preload("ShiftSchedule").First(&attendance, id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Attendance{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AttendanceRepo) Delete(id uint) error {
	return r.db.Delete(&model.Attendance{}, id).Error
}

func (r *AttendanceRepo) List(employeeID *uint, offset, limit int) ([]model.Attendance, int64, error) {
	var attendances []model.Attendance
	var total int64

	// Khởi tạo truy vấn cơ sở
	query := r.db.Model(&model.Attendance{})

	// Thêm điều kiện lọc nếu employeeID được cung cấp
	if employeeID != nil {
		// GORM cần JOIN với ShiftSchedule để truy cập EmployeeID
		query = query.
			Joins("JOIN shift_schedules ON shift_schedules.id = attendances.shift_schedule_id").
			Where("shift_schedules.employee_id = ?", *employeeID)
	}

	// 1. Đếm tổng số lượng bản ghi (trước khi áp dụng offset/limit)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 2. Thực hiện truy vấn lấy dữ liệu
	if err := query.
		// Preload các mối quan hệ (tôi thêm Preload cho Employee và Shift để dữ liệu đầy đủ hơn)
		Preload("ShiftSchedule").
		Preload("ShiftSchedule.Employee").
		Preload("ShiftSchedule.Shift").
		Offset(offset).
		Limit(limit).
		Find(&attendances).Error; err != nil {
		return nil, 0, err
	}

	return attendances, total, nil
}
