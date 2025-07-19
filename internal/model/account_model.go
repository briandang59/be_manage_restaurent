package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserName  string         `json:"user_name"`
	Password  string         `json:"password"`
	RoleId    uint           `json:"role_id"`
	Role      *Role           `json:"roles,omitempty" gorm:"foreignKey:RoleId"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"           gorm:"index"`
}
