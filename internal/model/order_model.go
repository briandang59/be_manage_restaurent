package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CustomerId uint           `json:"customer_id" gorm:"not null"`
	TableId    uint           `json:"table_id" gorm:"not null"`
	Amount     int64          `json:"amount" gorm:"not null"`
	Status     enum.OrderEnum `json:"status" gorm:"default:'UnPaid';not null"`
	Memo       string         `json:"memo"`
	Customer   *Customer      `json:"customer,omitempty" gorm:"foreignKey:CustomerId"`
	Table      *Table         `json:"table,omitempty" gorm:"foreignKey:TableId"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
