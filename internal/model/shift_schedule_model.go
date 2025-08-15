package model

import (
	"time"

	"gorm.io/gorm"
)

type ShiftSchedule struct {
	ID         uint           `json:"id"`
	EmployeeID *uint          `json:"employee_id" gorm:"not null"`
	ShiftID    *uint          `json:"shift_id" gorm:"not null"`
	Date       string         `json:"date" gorm:"not null"`
	Employee   *Employee      `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
	Shift      *Shift         `json:"shift,omitempty" gorm:"foreignKey:ShiftID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"           gorm:"index"`
}
