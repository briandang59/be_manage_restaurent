package dto

type CreateRoleDTO struct {
	RoleName      string `json:"role_name" binding:"required"`
	PermissionIDs []uint `json:"permissions"`
}

type UpdateRoleDTO struct {
	RoleName      string `json:"role_name,omitempty"`
	PermissionIDs []uint `json:"permissions,omitempty"`
}
