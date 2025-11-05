package model

import (
	"time"

	"gorm.io/gorm"
)

type Table struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	TableName  string         `json:"table_name" gorm:"not null"`
	Position   string         `json:"position"`
	Seats      int64          `json:"seats"`
	Memo       string         `json:"memo"`
	IsOccupied bool           `json:"is_occupied"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"           gorm:"index"`
}
