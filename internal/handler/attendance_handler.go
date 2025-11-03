package handler

import (
	"manage_restaurent/internal/dto"
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	svc *service.AttendanceService
}

func NewAttendanceHandler(s *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách chấm công
// @Description Lấy danh sách chấm công
// @Tags attendance
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Attendance}
// @Router /attendances [get]
func (h *AttendanceHandler) GetAll(c *gin.Context) {
	// Lấy và xử lý phân trang
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// Lấy employee_id từ query parameter
	employeeIDStr := c.Query("employee_id")
	var employeeID *uint // Sử dụng con trỏ để có thể là nil/null

	if employeeIDStr != "" {
		// Chuyển đổi string sang uint
		id, err := strconv.ParseUint(employeeIDStr, 10, 32)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid employee_id format")
			return
		}
		uID := uint(id)
		employeeID = &uID
	}

	// Gọi Service với tham số employeeID mới
	list, total, err := h.svc.List(employeeID, offset, pageSize)

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
// @Summary Lấy chi tiết chấm công
// @Description Lấy chi tiết một bản ghi chấm công theo ID
// @Tags attendance
// @Produce json
// @Param id path int true "ID chấm công"
// @Success 200 {object} model.Attendance
// @Failure 404 {object} response.ErrorResponse
// @Router /attendances/{id} [get]
func (h *AttendanceHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	attendance, err := h.svc.GetByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Attendance not found")
		return
	}
	response.Success(c, attendance, nil)
}

// Create godoc
// @Summary Tạo mới chấm công
// @Description Tạo mới một bản ghi chấm công
// @Tags attendance
// @Accept json
// @Produce json
// @Param attendance body dto.CreateAttendanceDTO true "Dữ liệu chấm công" example({"shift_schedule_id":1,"actual_start_time":"2024-07-22T08:00:00Z","actual_end_time":"2024-07-22T17:00:00Z"})
// @Success 200 {object} model.Attendance
// @Failure 400 {object} response.ErrorResponse
// @Router /attendances [post]
func (h *AttendanceHandler) Create(c *gin.Context) {
	var req dto.CreateAttendanceDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Chuyển đổi DTO thành model
	attendance := model.Attendance{
		ShiftScheduleId: req.ShiftScheduleId,
		ActualStartTime: req.ActualStartTime,
		ActualEndTime:   req.ActualEndTime,
	}

	if err := h.svc.Create(&attendance); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, attendance, nil)
}

// Update godoc
// @Summary Cập nhật chấm công
// @Description Cập nhật thông tin một bản ghi chấm công
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "ID chấm công"
// @Param updates body object true "Dữ liệu cập nhật" example({"actual_end_time":"2024-07-22T18:00:00Z"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /attendances/{id} [patch]
func (h *AttendanceHandler) Update(c *gin.Context) {
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
	response.Success(c, "Attendance updated successfully", nil)
}

// Delete godoc
// @Summary Xóa chấm công
// @Description Xóa một bản ghi chấm công
// @Tags attendance
// @Produce json
// @Param id path int true "ID chấm công"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /attendances/{id} [delete]
func (h *AttendanceHandler) Delete(c *gin.Context) {
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
	response.Success(c, "Attendance deleted successfully", nil)
}
