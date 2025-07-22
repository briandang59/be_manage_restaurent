package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách order
// @Description Lấy danh sách đơn hàng
// @Tags order
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Order}
// @Router /orders [get]
func (h *OrderHandler) GetAll(c *gin.Context) {
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
// @Summary Lấy chi tiết order
// @Description Lấy chi tiết một đơn hàng theo ID
// @Tags order
// @Produce json
// @Param id path int true "ID order"
// @Success 200 {object} model.Order
// @Failure 404 {object} response.ErrorResponse
// @Router /orders/{id} [get]
func (h *OrderHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	order, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Order not found")
		return
	}
	response.Success(c, order, nil)
}

// Create godoc
// @Summary Tạo mới order
// @Description Tạo mới một đơn hàng
// @Tags order
// @Accept json
// @Produce json
// @Param order body model.Order true "Dữ liệu order" example({"customer_id":1,"table_id":2,"amount":500000,"status":"UnPaid"})
// @Success 200 {object} model.Order
// @Failure 400 {object} response.ErrorResponse
// @Router /orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&order); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, order, nil)
}

// Update godoc
// @Summary Cập nhật order
// @Description Cập nhật thông tin đơn hàng
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "ID order"
// @Param updates body object true "Dữ liệu cập nhật" example({"status":"Paid"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /orders/{id} [patch]
func (h *OrderHandler) Update(c *gin.Context) {
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
	response.Success(c, "Order updated successfully", nil)
}

// Delete godoc
// @Summary Xóa order
// @Description Xóa một đơn hàng
// @Tags order
// @Produce json
// @Param id path int true "ID order"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /orders/{id} [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Order deleted successfully", nil)
} 