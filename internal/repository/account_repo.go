package repository

import (
	"gorm.io/gorm"
	"manage_restaurent/internal/model"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepo) GetByID(id uint) (*model.Account, error) {
	var account model.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) GetByUserName(username string) (*model.Account, error) {
	var account model.Account
	if err := r.db.Where("user_name = ?", username).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.Account{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AccountRepo) Delete(id uint) error {
	return r.db.Delete(&model.Account{}, id).Error
}

func (r *AccountRepo) List(offset, limit int) ([]model.Account, int64, error) {
	var accounts []model.Account
	var total int64
	r.db.Model(&model.Account{}).Count(&total)
	if err := r.db.Offset(offset).Limit(limit).Find(&accounts).Error; err != nil {
		return nil, 0, err
	}
	return accounts, total, nil
} 