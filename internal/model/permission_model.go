package model

type Permission struct {
	ID    uint    `json:"id"`
	Name  string  `json:"permission_name" gorm:"type:text;not null;unique"`
	Roles *[]Role `json:"roles" gorm:"many2many:role_permissions;"`
}
