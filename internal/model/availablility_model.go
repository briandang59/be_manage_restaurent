package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Availibility struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	EmployeeId  uint           `json:"employee_id"`
	ShiftId     uint           `json:"shift_id" gorm:"not null"`
	DayOfWeek   enum.DayOfWeek `json:"day_of_week" gorm:"type:varchar(10);not null"`
	IsAvailable bool           `json:"is_available" gorm:"default:false"`
	Employee    Employee       `json:"employee" gorm:"foreignKey:EmployeeID"`
	Shift       Shift          `json:"shifts" gorm:"foreignKey:ShiftID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}
