package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID         uint               `json:"id" gorm:"primaryKey"`
	OrderId    *uint              `json:"order_id" gorm:"not null"`
	MenuItemId *uint              `json:"menu_item_id" gorm:"not null"`
	Quantity   int64              `json:"quantity" gorm:"not null"`
	Amount     int64              `json:"amount"`
	Memo       string             `json:"memo"`
	Status     enum.OrderItemEnum `json:"status" gorm:"default:Pending"`
	Order      *Order             `json:"order,omitempty" gorm:"foreignKey:OrderId"`
	MenuItem   *MenuItem          `json:"menu_item,omitempty" gorm:"foreignKey:MenuItemId"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DeletedAt  gorm.DeletedAt     `json:"-" gorm:"index"`
}
