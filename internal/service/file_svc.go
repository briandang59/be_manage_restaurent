package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type FileService struct {
	repo *repository.FileRepo
}

func NewFileService(r *repository.FileRepo) *FileService {
	return &FileService{repo: r}
}

func (s *FileService) Create(file *model.File) error {
	return s.repo.Create(file)
}

func (s *FileService) GetByID(id uint) (*model.File, error) {
	return s.repo.GetByID(id)
}

func (s *FileService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *FileService) List(offset, limit int) ([]model.File, int64, error) {
	return s.repo.List(offset, limit)
} 