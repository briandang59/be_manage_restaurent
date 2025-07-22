package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	svc *service.TicketService
}

func NewTicketHandler(s *service.TicketService) *TicketHandler {
	return &TicketHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách ticket
// @Description Lấy danh sách phiếu nhập/xuất kho
// @Tags ticket
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Ticket}
// @Router /tickets [get]
func (h *TicketHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	list, total, err := h.svc.List(offset, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list, &response.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
	})
}

// GetByID godoc
// @Summary Lấy chi tiết ticket
// @Description Lấy chi tiết một phiếu theo ID
// @Tags ticket
// @Produce json
// @Param id path int true "ID ticket"
// @Success 200 {object} model.Ticket
// @Failure 404 {object} response.ErrorResponse
// @Router /tickets/{id} [get]
func (h *TicketHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	ticket, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Ticket not found")
		return
	}
	response.Success(c, ticket, nil)
}

// Create godoc
// @Summary Tạo mới ticket
// @Description Tạo mới một phiếu nhập/xuất kho
// @Tags ticket
// @Accept json
// @Produce json
// @Param ticket body model.Ticket true "Dữ liệu ticket" example({"ingredient_id":1,"quantity":10,"unit":"kg","ticket_type":"IMPORT"})
// @Success 200 {object} model.Ticket
// @Failure 400 {object} response.ErrorResponse
// @Router /tickets [post]
func (h *TicketHandler) Create(c *gin.Context) {
	var ticket model.Ticker
	if err := c.ShouldBindJSON(&ticket); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&ticket); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, ticket, nil)
}

// Update godoc
// @Summary Cập nhật ticket
// @Description Cập nhật thông tin một phiếu
// @Tags ticket
// @Accept json
// @Produce json
// @Param id path int true "ID ticket"
// @Param updates body object true "Dữ liệu cập nhật" example({"quantity":20})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /tickets/{id} [patch]
func (h *TicketHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Update(uint(id), updates); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, "Ticket updated successfully", nil)
}

// Delete godoc
// @Summary Xóa ticket
// @Description Xóa một phiếu
// @Tags ticket
// @Produce json
// @Param id path int true "ID ticket"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /tickets/{id} [delete]
func (h *TicketHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, "Ticket deleted successfully", nil)
} 