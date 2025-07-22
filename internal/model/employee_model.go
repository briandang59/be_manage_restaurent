package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID             uint                      `json:"id" gorm:"primaryKey"`
	FullName       string                    `json:"full_name"`
	Gender         bool                      `json:"gender"`
	Birthday       string                    `json:"birthday"`
	PhoneNumber    string                    `json:"phone_number"`
	Email          string                    `json:"email"`
	ScheduleType   enum.EmployeeScheduleType `json:"schedule_type"`
	Address        string                    `json:"address"`
	JoinDate       string                    `json:"join_date"`
	BaseSalary     int64                     `json:"base_salary"`
	SalaryPerHours int64                     `json:"salary_per_hour"`
	AccountID      *uint                     `json:"account_id"`
	Account        *Account                  `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	AvatarFileID   *uint                     `json:"avatar_file_id"`
	AvatarFile     *File                     `json:"avatar_file,omitempty" gorm:"foreignKey:AvatarFileID"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	DeletedAt      gorm.DeletedAt            `json:"-" gorm:"index"`
}
