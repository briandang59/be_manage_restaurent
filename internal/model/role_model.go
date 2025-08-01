package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	RoleName    string         `json:"role_name" gorm:"unique;not null"`
	Permissions *[]Permission  `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
