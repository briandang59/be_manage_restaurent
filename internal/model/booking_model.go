package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CustomerName string         `json:"customer_name" gorm:"not null"`
	PhoneNumber  string         `json:"phone_number" gorm:"not null"`
	Email        string         `json:"email" gorm:"not null"`
	BookingDate  string         `json:"booking_date" gorm:"not null"`
	BookingTime  string         `json:"booking_time" gorm:"not null"`
	TotalPersons int            `json:"total_persons" gorm:"not null"`
	Status       string         `json:"status" gorm:"not null;default:'pending'"`
	Memo         string         `json:"memo" gorm:"type:text;default:null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"           gorm:"index"`
}
