package dto

import (
	"manage_restaurent/internal/model"
	"time"
)

type AccountDTO struct {
	ID        uint        `json:"id"`
	UserName  string      `json:"user_name"`
	RoleId    uint        `json:"role_id"`
	Role      *model.Role `json:"role,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type LoginRequestDTO struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
	Role  interface{} `json:"role"`
}
