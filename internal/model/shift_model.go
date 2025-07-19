package model

import (
	"time"

	"gorm.io/gorm"
)

type Shift struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ShiftName string         `json:"shift_name" gorm:"not null"`
	Code      string         `json:"code" gorm:"not null"`
	StartTime string         `json:"start_time" gorm:"not null"`
	EndTime   string         `json:"end_time" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
