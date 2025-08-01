package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"permission_name" gorm:"type:text;not null;unique"`
	Roles     *[]Role        `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
