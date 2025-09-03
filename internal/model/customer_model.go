package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	FullName    string         `json:"full_name"`
	PhoneNumber string         `json:"phone_number" gorm:"uniqueIndex"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}
