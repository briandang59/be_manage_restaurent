package dto

import (
	"manage_restaurent/internal/enum"
)

// CreateEmployeeDTO DTO để tạo employee với account tự động
type CreateEmployeeDTO struct {
	FullName       string                    `json:"full_name" binding:"required"`
	Gender         bool                      `json:"gender"`
	Birthday       string                    `json:"birthday"`
	PhoneNumber    string                    `json:"phone_number" binding:"required"`
	Email          string                    `json:"email" binding:"required,email"`
	ScheduleType   enum.EmployeeScheduleType `json:"schedule_type"`
	Address        string                    `json:"address"`
	JoinDate       string                    `json:"join_date"`
	BaseSalary     int64                     `json:"base_salary"`
	SalaryPerHour  int64                     `json:"salary_per_hour"`
	AvatarFileID   *uint                     `json:"avatar_file_id"`
	RoleId         uint                      `json:"role_id" binding:"required"`
} 