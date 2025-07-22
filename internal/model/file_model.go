package model

import (
	"time"
	"gorm.io/gorm"
)

type File struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	FileName  string         `json:"file_name" gorm:"not null"`
	Url       string         `json:"url" gorm:"not null"`
	MimeType  string         `json:"mime_type"`
	Size      int64          `json:"size"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
} 