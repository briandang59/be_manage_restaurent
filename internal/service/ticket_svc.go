package service

import (
	"fmt"
	"os"

	"manage_restaurent/internal/model"
	"manage_restaurent/internal/repository"
	"manage_restaurent/utils"
)

type TicketService struct {
	repo            *repository.TicketRepo
	ingredientRepo  *repository.IngredientRepo
	telegramService *TelegramService
}

func NewTicketService(r *repository.TicketRepo, ingredientRepo *repository.IngredientRepo, telegramSvc *TelegramService) *TicketService {
	return &TicketService{
		repo:            r,
		ingredientRepo:  ingredientRepo,
		telegramService: telegramSvc,
	}
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
	if ticket.TicketType == "Export" {
		if ingredient.Quantity == 0 {
			return fmt.Errorf("đã hết hàng")
		}
		if ticket.Quantity > ingredient.Quantity {
			return fmt.Errorf("không đủ số lượng trong kho: chỉ còn %d %s", ingredient.Quantity, ingredient.Unit)
		}
	}
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

	if err := s.ingredientRepo.Update(ticket.IngredientId, updates); err != nil {
		return err
	}

	if newQuantity <= ingredient.WarningQuantity {
		// Gửi thông báo Telegram
		botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
		chatID := utils.TelegramChatID
		if botToken == "" || chatID == "" {
			// Log error nếu cần, nhưng không throw để tránh ảnh hưởng business logic
			fmt.Printf("Telegram config missing: BOT_TOKEN or CHAT_ID not set\n")
			return nil
		}
		message := fmt.Sprintf("Cảnh báo: Nguyên liệu '%s' chỉ còn %d %s. Giới hạn cảnh báo: %d %s.",
			ingredient.Name,
			newQuantity,
			ticket.Unit,
			ingredient.WarningQuantity,
			ticket.Unit,
		)
		req := model.TelegramSendRequest{
			BotToken: botToken,
			ChatID:   chatID,
			Message:  message,
		}
		_, err := s.telegramService.SendMessage(req)
		if err != nil {
			// Không throw error, chỉ log nếu cần, để không ảnh hưởng business logic
			fmt.Printf("Failed to send Telegram notification: %v\n", err)
		}
	}

	return nil
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
