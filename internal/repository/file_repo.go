package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type FileRepo struct {
	db *gorm.DB
}

func NewFileRepo(db *gorm.DB) *FileRepo {
	return &FileRepo{db: db}
}

func (r *FileRepo) Create(file *model.File) error {
	return r.db.Create(file).Error
}

func (r *FileRepo) GetByID(id uint) (*model.File, error) {
	var file model.File
	if err := r.db.First(&file, id).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *FileRepo) Delete(id uint) error {
	return r.db.Delete(&model.File{}, id).Error
}

func (r *FileRepo) List(offset, limit int, preloadFields []string) ([]model.File, int64, error) {
	var files []model.File
	var total int64
	
	query := r.db.Model(&model.File{})
	for _, field := range preloadFields {
		query = query.Preload(field)
	}
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	if err := query.Offset(offset).Limit(limit).Find(&files).Error; err != nil {
		return nil, 0, err
	}
	return files, total, nil
} 