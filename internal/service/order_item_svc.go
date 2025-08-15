// internal/service/order_item_service.go
package service

import (
	"fmt"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type OrderItemService struct {
	repo *repository.OrderItemRepo
}

func NewOrderItemService(r *repository.OrderItemRepo) *OrderItemService {
	return &OrderItemService{repo: r}
}

// CREATE: amount = menu_item.price * quantity, đồng thời cộng vào Order.amount
func (s *OrderItemService) Create(orderItem *model.OrderItem) error {
	if orderItem.OrderId == nil || orderItem.MenuItemId == nil {
		return fmt.Errorf("order_id và menu_item_id là bắt buộc")
	}
	if orderItem.Quantity <= 0 {
		return fmt.Errorf("quantity phải > 0")
	}

	return s.repo.ExecTx(func(txRepo *repository.OrderItemRepo) error {
		price, err := txRepo.GetMenuItemPrice(*orderItem.MenuItemId)
		if err != nil {
			return err
		}

		orderItem.Amount = price * orderItem.Quantity

		if err := txRepo.Create(orderItem); err != nil {
			return err
		}

		return txRepo.AddOrderAmount(*orderItem.OrderId, orderItem.Amount)
	})
}

// UPDATE: nếu đổi quantity/menu_item_id thì tính lại amount và cập nhật chênh lệch vào Order
func (s *OrderItemService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.ExecTx(func(txRepo *repository.OrderItemRepo) error {
		// Lấy bản hiện tại
		cur, err := txRepo.GetByID(id)
		if err != nil {
			return err
		}
		if cur.OrderId == nil {
			return fmt.Errorf("order_item không gắn order")
		}
		oldAmount := cur.Amount

		// Tính ra quantity & menuItemId mới
		newQty := cur.Quantity
		if v, ok := updates["quantity"]; ok {
			switch t := v.(type) {
			case int:
				newQty = int64(t)
			case int64:
				newQty = t
			case float64:
				newQty = int64(t)
			default:
				return fmt.Errorf("quantity không hợp lệ")
			}
		}
		newMenuID := cur.MenuItemId
		if v, ok := updates["menu_item_id"]; ok {
			switch t := v.(type) {
			case int:
				u := uint(t)
				newMenuID = &u
			case int64:
				u := uint(t)
				newMenuID = &u
			case float64:
				u := uint(t)
				newMenuID = &u
			default:
				return fmt.Errorf("menu_item_id không hợp lệ")
			}
		}
		if newMenuID == nil {
			return fmt.Errorf("menu_item_id là bắt buộc")
		}
		if newQty <= 0 {
			return fmt.Errorf("quantity phải > 0")
		}

		// Tính amount mới
		price, err := txRepo.GetMenuItemPrice(*newMenuID)
		if err != nil {
			return err
		}
		newAmount := price * newQty

		// Ghi amount mới vào updates để lưu DB
		updates["amount"] = newAmount

		// Cập nhật OrderItem
		if err := txRepo.Update(id, updates); err != nil {
			return err
		}

		// Cập nhật chênh lệch vào tổng Order
		delta := newAmount - oldAmount
		if delta != 0 {
			if err := txRepo.AddOrderAmount(*cur.OrderId, delta); err != nil {
				return err
			}
		}
		return nil
	})
}

// DELETE: xóa item và trừ amount khỏi tổng Order
func (s *OrderItemService) Delete(id uint) error {
	return s.repo.ExecTx(func(txRepo *repository.OrderItemRepo) error {
		cur, err := txRepo.GetByID(id)
		if err != nil {
			return err
		}
		if err := txRepo.Delete(id); err != nil {
			return err
		}
		if cur.OrderId != nil && cur.Amount != 0 {
			return txRepo.AddOrderAmount(*cur.OrderId, -cur.Amount)
		}
		return nil
	})
}

func (s *OrderItemService) GetByID(id uint) (*model.OrderItem, error) {
	return s.repo.GetByID(id)
}

func (s *OrderItemService) ListByOrderID(orderID uint, offset, limit int) ([]model.OrderItem, int64, error) {
	return s.repo.ListByOrderID(orderID, offset, limit)
}

// internal/service/order_item_service.go
func (s *OrderItemService) List(offset, limit int) ([]model.OrderItem, int64, error) {
	return s.repo.List(offset, limit)
}
