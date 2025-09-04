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

func (r *TicketRepo) Create(ticket *model.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *TicketRepo) GetByID(id uint, preloadFields []string) (*model.Ticket, error) {
	var ticket model.Ticket
	db := r.db.Model(&model.Ticket{})
	for _, field := range preloadFields {
		db = db.Preload(field)
	}
	if err := db.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Ticket{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TicketRepo) Delete(id uint) error {
	return r.db.Delete(&model.Ticket{}, id).Error
}

func (r *TicketRepo) List(offset, limit int, preloadFields []string) ([]model.Ticket, int64, error) {
	var tickets []model.Ticket
	var total int64
	db := r.db.Model(&model.Ticket{})
	db.Count(&total)
	for _, field := range preloadFields {
		db = db.Preload(field)
	}
	if err := db.Offset(offset).Limit(limit).Find(&tickets).Error; err != nil {
		return nil, 0, err
	}
	return tickets, total, nil
}
