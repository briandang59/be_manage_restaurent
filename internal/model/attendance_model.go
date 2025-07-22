package model

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	ShiftScheduleId uint           `json:"shift_schedule_id"`
	ActualStartTime time.Time      `json:"actual_start_time"`
	ActualEndTime   time.Time      `json:"actual_end_time"`
	Hours           int64          `json:"hours"`
	ShiftSchedule   *ShiftSchedule `json:"shift_schedule" gorm:"foreignKey:ShiftScheduleId"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}
