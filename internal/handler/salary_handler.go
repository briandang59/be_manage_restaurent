package handler

import (
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SalaryHandler struct {
	svc *service.SalaryService
}

func NewSalaryHandler(s *service.SalaryService) *SalaryHandler {
	return &SalaryHandler{svc: s}
}

// GetSalaryByEmployeeAndMonth godoc
// @Summary Tính lương nhân viên theo tháng
// @Description Tính lương cho một nhân viên theo tháng cụ thể
// @Tags salary
// @Produce json
// @Param employee_id path int true "ID nhân viên"
// @Param month query string false "Tháng (YYYY-MM)" example("2024-09")
// @Success 200 {object} response.Body{data=map[string]interface{}}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /salary/employees/{employee_id} [get]
func (h *SalaryHandler) GetSalaryByEmployeeAndMonth(c *gin.Context) {
	employeeIDStr := c.Param("employee_id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid employee ID")
		return
	}
	month := c.Query("month")
	if month == "" {
		response.Error(c, http.StatusBadRequest, "Month is required (format: YYYY-MM)")
		return
	}
	salaryData, err := h.svc.CalculateSalary(uint(employeeID), month)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, salaryData, nil)
}

// GetAllSalaries godoc
// @Summary Tính lương toàn bộ nhân viên theo tháng
// @Description Tính lương cho tất cả nhân viên trong tháng cụ thể
// @Tags salary
// @Produce json
// @Param month query string true "Tháng (YYYY-MM)" example("2024-09")
// @Success 200 {object} response.Body{data=[]map[string]interface{}}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /salary/all [get]
func (h *SalaryHandler) GetAllSalaries(c *gin.Context) {
	month := c.Query("month")
	if month == "" {
		response.Error(c, http.StatusBadRequest, "Month is required (format: YYYY-MM)")
		return
	}
	salaries, err := h.svc.CalculateAllSalaries(month)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, salaries, nil)
}
