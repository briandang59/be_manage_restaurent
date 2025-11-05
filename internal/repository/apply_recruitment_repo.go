package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type ApplyRecruitmentRepo struct {
	db *gorm.DB
}

func NewApplyRecruitmentRepo(db *gorm.DB) *ApplyRecruitmentRepo {
	return &ApplyRecruitmentRepo{db: db}
}

func (r *ApplyRecruitmentRepo) Create(apply_recruitment *model.ApplyRecruitment) error {
	return r.db.Create(apply_recruitment).Error
}

func (r *ApplyRecruitmentRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.ApplyRecruitment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *ApplyRecruitmentRepo) Delete(id uint) error {
	return r.db.Delete(&model.ApplyRecruitment{}, id).Error
}

func (r *ApplyRecruitmentRepo) FindAll(page, pageSize int, preloadFields []string) ([]model.ApplyRecruitment, int64, error) {
	var list []model.ApplyRecruitment
	var total int64
	offset := (page - 1) * pageSize

	query := r.db.Model(&model.ApplyRecruitment{})

	for _, field := range preloadFields {
		query = query.Preload(field)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
