package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type MenuItem struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       int64          `json:"price" gorm:"not null"`
	FileID      *uint          `json:"file_id"`
	File        *File          `json:"file,omitempty" gorm:"foreignKey:FileID"`
	Status      enum.MenuItem  `json:"status" gorm:"default:'Available';not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
