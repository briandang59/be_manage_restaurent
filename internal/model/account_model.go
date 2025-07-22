package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserName  string         `json:"user_name" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	RoleId    uint           `json:"role_id" gorm:"not null"`
	Role      *Role          `json:"role,omitempty" gorm:"foreignKey:RoleId"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
