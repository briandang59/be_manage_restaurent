package model

import (
	"time"

	"gorm.io/gorm"
)

type Recruitment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"type:text;not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	IsOpen    bool           `json:"is_open" gorm:"type:boolean;default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
