package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"category_name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
