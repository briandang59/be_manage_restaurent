package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Ticker struct {
	ID           uint            `json:"id" gorm:"primaryKey"`
	IngredientId uint            `json:"ingredient_id"`
	Quantity     int64           `json:"quantity"`
	Unit         int64           `json:"unit"`
	TicketType   enum.TicketType `json:"ticket_type"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"-"           gorm:"index"`
}
