package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	OrderId    uint           `json:"order_id" gorm:"not null"`
	MenuItemId uint           `json:"menu_item_id" gorm:"not null"`
	Quantity   int64          `json:"quantity"`
	Amount     int64          `json:"amount" gorm:"not null"`
	Memo       string         `json:"memo"`
	Order      *Order         `json:"order,omitempty" gorm:"foreignKey:OrderId"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
