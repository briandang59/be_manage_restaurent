package handler

import (
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

func (h *AttendanceHandler) GetAll(c *gin.Context) {
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

func (h *AttendanceHandler) Create(c *gin.Context) {
	var attendance model.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.Create(&attendance); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, attendance, nil)
}

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