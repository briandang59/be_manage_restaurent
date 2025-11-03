package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type RecruitmentRepo struct {
	db *gorm.DB
}

func NewRecruitmentRepo(db *gorm.DB) *RecruitmentRepo {
	return &RecruitmentRepo{db: db}
}

func (r *RecruitmentRepo) Create(recruitment *model.Recruitment) error {
	return r.db.Create(recruitment).Error
}

func (r *RecruitmentRepo) Delete(id uint) error {
	return r.db.Delete(&model.Recruitment{}, id).Error
}

func (r *RecruitmentRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Recruitment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RecruitmentRepo) List(offset, limit int) ([]model.Recruitment, int64, error) {
	var recruitments []model.Recruitment
	var total int64
	r.db.Model(&model.Recruitment{}).Count(&total)
	if err := r.db.Offset(offset).Limit(limit).Find(&recruitments).Error; err != nil {
		return nil, 0, err
	}
	return recruitments, total, nil
}
