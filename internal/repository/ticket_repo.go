package repository

import (
	"manage_restaurent/internal/model"
	"gorm.io/gorm"
)

type TicketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepo {
	return &TicketRepo{db: db}
}

func (r *TicketRepo) Create(ticket *model.Ticker) error {
	return r.db.Create(ticket).Error
}

func (r *TicketRepo) GetByID(id uint) (*model.Ticker, error) {
	var ticket model.Ticker
	if err := r.db.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Ticker{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TicketRepo) Delete(id uint) error {
	return r.db.Delete(&model.Ticker{}, id).Error
}

func (r *TicketRepo) List(offset, limit int) ([]model.Ticker, int64, error) {
	var tickets []model.Ticker
	var total int64
	r.db.Model(&model.Ticker{}).Count(&total)
	if err := r.db.Offset(offset).Limit(limit).Find(&tickets).Error; err != nil {
		return nil, 0, err
	}
	return tickets, total, nil
} 