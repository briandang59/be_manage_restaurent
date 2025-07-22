package handler

import (
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"manage_restaurent/internal/model"

	"github.com/gin-gonic/gin"
)

// ShiftHandler xử lý các request HTTP cho Shift
type ShiftHandler struct {
	svc *service.ShiftService
}

// NewShiftHandler tạo một thể hiện mới của ShiftHandler
func NewShiftHandler(s *service.ShiftService) *ShiftHandler {
	return &ShiftHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách ca làm việc
// @Description Lấy danh sách ca làm việc
// @Tags shift
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Shift}
// @Router /shifts [get]
func (h *ShiftHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	list, total, err := h.svc.GetAll(page, pageSize)
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
// @Summary Lấy chi tiết ca làm việc
// @Description Lấy chi tiết một ca làm việc theo ID
// @Tags shift
// @Produce json
// @Param id path int true "ID shift"
// @Success 200 {object} model.Shift
// @Failure 404 {object} response.ErrorResponse
// @Router /shifts/{id} [get]
func (h *ShiftHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	shift, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Shift not found")
		return
	}

	response.Success(c, shift, nil)
}

// Create godoc
// @Summary Tạo mới ca làm việc
// @Description Tạo mới một ca làm việc
// @Tags shift
// @Accept json
// @Produce json
// @Param shift body model.Shift true "Dữ liệu ca làm việc" example({"shift_name":"Ca sáng","code":"MORNING","start_time":"08:00","end_time":"12:00"})
// @Success 200 {object} model.Shift
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts [post]
func (h *ShiftHandler) Create(c *gin.Context) {
	var shift model.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Create(&shift); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, shift, nil)
}

// Update godoc
// @Summary Cập nhật ca làm việc
// @Description Cập nhật thông tin ca làm việc
// @Tags shift
// @Accept json
// @Produce json
// @Param id path int true "ID shift"
// @Param updates body object true "Dữ liệu cập nhật" example({"shift_name":"Ca chiều"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts/{id} [patch]
func (h *ShiftHandler) Update(c *gin.Context) {
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

	response.Success(c, "Shift updated successfully", nil)
}

// Delete godoc
// @Summary Xóa ca làm việc
// @Description Xóa một ca làm việc
// @Tags shift
// @Produce json
// @Param id path int true "ID shift"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /shifts/{id} [delete]
func (h *ShiftHandler) Delete(c *gin.Context) {
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

	response.Success(c, "Shift deleted successfully", nil)
}
