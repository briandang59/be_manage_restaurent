package model

import (
	"time"

	"gorm.io/gorm"
)

type Ingredient struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Quantity        int64          `json:"quantity"`
	WarningQuantity int64          `json:"warning_quantity" gorm:"not null"`
	Supplier        string         `json:"supplier"`
	Unit            string         `json:"unit"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-"           gorm:"index"`
}
