package model

import (
	"time"

	"gorm.io/gorm"
)

type ApplyRecruitment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	RecruitmentID uint           `json:"recruitment_id"`
	FullName      string         `json:"fullname"`
	Email         string         `json:"email"`
	PhoneNumber   string         `json:"phone_number"`
	CVID          uint           `json:"cv_id"`
	Recruitment   *Recruitment   `json:"recruitment" gorm:"foreignKey:RecruitmentID"`
	CV            *File          `json:"cv" gorm:"foreignKey:CVID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
