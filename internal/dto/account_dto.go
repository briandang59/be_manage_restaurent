package dto

import (
	"time"
	"manage_restaurent/internal/model"
)

type AccountDTO struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	RoleId    uint      `json:"role_id"`
	Role      *model.Role     `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} 