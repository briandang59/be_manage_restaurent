// internal/repository/order_item_repo.go
package repository

import (
	"manage_restaurent/internal/model"

	"gorm.io/gorm"
)

type OrderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) *OrderItemRepo {
	return &OrderItemRepo{db: db}
}

// Chạy 1 hàm trong transaction; cấp 1 repo gắn với tx
func (r *OrderItemRepo) ExecTx(fn func(txRepo *OrderItemRepo) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&OrderItemRepo{db: tx})
	})
}

func (r *OrderItemRepo) Create(orderItem *model.OrderItem) error {
	return r.db.Create(orderItem).Error
}

func (r *OrderItemRepo) GetByID(id uint) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	// preload luôn cả MenuItem để dùng khi cần
	if err := r.db.Preload("Order").Preload("MenuItem").First(&orderItem, id).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemRepo) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&model.OrderItem{}).Where("id = ?", id).Updates(updates).Error
}

func (r *OrderItemRepo) Delete(id uint) error {
	return r.db.Delete(&model.OrderItem{}, id).Error
}

// internal/repository/order_item_repo.go

func (r *OrderItemRepo) ListByOrderID(orderID uint, offset, limit int) ([]model.OrderItem, int64, error) {
	var items []model.OrderItem
	var total int64

	base := r.db.Model(&model.OrderItem{}).Where("order_id = ?", orderID)
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.
		Preload("Order").
		Preload("MenuItem").
		Where("order_id = ?", orderID).
		Offset(offset).
		Limit(limit).
		Order("id DESC").
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// === Helpers cho nghiệp vụ ===

// Lấy price của MenuItem (báo lỗi nếu không tồn tại)
func (r *OrderItemRepo) GetMenuItemPrice(menuItemID uint) (int64, error) {
	var mi model.MenuItem
	if err := r.db.Select("id", "price").First(&mi, menuItemID).Error; err != nil {
		return 0, err
	}
	return mi.Price, nil
}

// Cộng/trừ amount vào tổng Order (delta có thể âm)
func (r *OrderItemRepo) AddOrderAmount(orderID uint, delta int64) error {
	return r.db.Model(&model.Order{}).
		Where("id = ?", orderID).
		UpdateColumn("amount", gorm.Expr("COALESCE(amount,0) + ?", delta)).Error
}

// internal/repository/order_item_repo.go
func (r *OrderItemRepo) List(offset, limit int) ([]model.OrderItem, int64, error) {
	var items []model.OrderItem
	var total int64

	// Đếm tổng
	if err := r.db.Model(&model.OrderItem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Lấy danh sách (preload Order & MenuItem), sắp xếp ổn định
	if err := r.db.Preload("Order").Preload("MenuItem").
		Order("id DESC").
		Offset(offset).Limit(limit).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
