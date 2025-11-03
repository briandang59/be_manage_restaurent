package model

import (
	"manage_restaurent/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID           uint            `json:"id" gorm:"primaryKey"`
	IngredientId uint            `json:"ingredient_id" gorm:"not null"`
	Quantity     int64           `json:"quantity"`
	Unit         string          `json:"unit" gorm:"not null"`
	TicketType   enum.TicketType `json:"ticket_type"`
	Ingredient   *Ingredient     `json:"ingredient,omitempty" gorm:"foreignKey:IngredientId"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"-" gorm:"index"`
}
