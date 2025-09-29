package service

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
)

type TicketService struct {
	repo           *repository.TicketRepo
	ingredientRepo *repository.IngredientRepo
}

func NewTicketService(r *repository.TicketRepo, ingredientRepo *repository.IngredientRepo) *TicketService {
	return &TicketService{repo: r, ingredientRepo: ingredientRepo}
}

func (s *TicketService) Create(ticket *model.Ticket) error {
	// Kiểm tra ingredient tồn tại
	ingredient, err := s.ingredientRepo.GetByID(ticket.IngredientId)
	if err != nil {
		return err
	}

	// Kiểm tra unit consistency (có thể bỏ qua nếu muốn linh hoạt)
	// if ingredient.Unit != ticket.Unit {
	// 	return fmt.Errorf("unit mismatch: ingredient unit is %s but ticket unit is %s", ingredient.Unit, ticket.Unit)
	// }

	// Tạo ticket trước
	if err := s.repo.Create(ticket); err != nil {
		return err
	}

	// Tính toán quantity mới dựa trên ticket type
	var newQuantity int64
	switch ticket.TicketType {
	case "Import":
		newQuantity = ingredient.Quantity + ticket.Quantity
	case "Export":
		newQuantity = ingredient.Quantity - ticket.Quantity
		if newQuantity < 0 {
			newQuantity = 0 // Không cho phép quantity âm
		}
	default:
		// Nếu không phải Import/Export thì không cập nhật
		return nil
	}

	// Cập nhật quantity của ingredient
	updates := map[string]interface{}{
		"quantity": newQuantity,
	}
	return s.ingredientRepo.Update(ticket.IngredientId, updates)
}

func (s *TicketService) GetByID(id uint, preloadFields []string) (*model.Ticket, error) {
	return s.repo.GetByID(id, preloadFields)
}

func (s *TicketService) Update(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}

func (s *TicketService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *TicketService) List(offset, limit int, preloadFields []string) ([]model.Ticket, int64, error) {
	return s.repo.List(offset, limit, preloadFields)
}
