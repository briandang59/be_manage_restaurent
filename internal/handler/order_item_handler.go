package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderItemHandler struct {
	svc *service.OrderItemService
}

func NewOrderItemHandler(s *service.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách order item
// @Description Lấy danh sách các món trong đơn hàng
// @Tags orderitem
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.OrderItem}
// @Router /order-items [get]
func (h *OrderItemHandler) GetAll(c *gin.Context) {
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
// @Summary Lấy chi tiết order item
// @Description Lấy chi tiết một món trong đơn hàng theo ID
// @Tags orderitem
// @Produce json
// @Param id path int true "ID order item"
// @Success 200 {object} model.OrderItem
// @Failure 404 {object} response.ErrorResponse
// @Router /order-items/{id} [get]
func (h *OrderItemHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	orderItem, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "OrderItem not found")
		return
	}
	response.Success(c, orderItem, nil)
}

// Create godoc
// @Summary Tạo mới order item
// @Description Thêm món vào đơn hàng
// @Tags orderitem
// @Accept json
// @Produce json
// @Param orderItem body model.OrderItem true "Dữ liệu order item" example({"order_id":1,"menu_item_id":2,"quantity":3,"amount":360000})
// @Success 200 {object} model.OrderItem
// @Failure 400 {object} response.ErrorResponse
// @Router /order-items [post]
func (h *OrderItemHandler) Create(c *gin.Context) {
	var orderItem model.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&orderItem); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, orderItem, nil)
}

// Update godoc
// @Summary Cập nhật order item
// @Description Cập nhật thông tin món trong đơn hàng
// @Tags orderitem
// @Accept json
// @Produce json
// @Param id path int true "ID order item"
// @Param updates body object true "Dữ liệu cập nhật" example({"quantity":5})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /order-items/{id} [patch]
func (h *OrderItemHandler) Update(c *gin.Context) {
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
	response.Success(c, "OrderItem updated successfully", nil)
}

// Delete godoc
// @Summary Xóa order item
// @Description Xóa một món khỏi đơn hàng
// @Tags orderitem
// @Produce json
// @Param id path int true "ID order item"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /order-items/{id} [delete]
func (h *OrderItemHandler) Delete(c *gin.Context) {
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
	response.Success(c, "OrderItem deleted successfully", nil)
} 