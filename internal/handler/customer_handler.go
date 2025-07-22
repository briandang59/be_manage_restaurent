package handler

import (
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"manage_restaurent/utils"
	"net/http"
	"strconv"

	"manage_restaurent/internal/model"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	svc *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách khách hàng
// @Description Lấy danh sách khách hàng
// @Tags customer
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Customer}
// @Router /customers [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	preloadFields := utils.ParsePopulateQuery(c.Request.URL.Query())

	list, total, err := h.svc.GetAll(page, pageSize, preloadFields)
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

// Create godoc
// @Summary Tạo mới khách hàng
// @Description Tạo mới một khách hàng
// @Tags customer
// @Accept json
// @Produce json
// @Param customer body model.Customer true "Dữ liệu khách hàng" example({"full_name":"Nguyễn Văn C","phone_number":"0123456789"})
// @Success 200 {object} model.Customer
// @Failure 400 {object} response.ErrorResponse
// @Router /customers [post]
func (h *CustomerHandler) Create(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Create(&customer); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, customer, nil)
}

// Update godoc
// @Summary Cập nhật khách hàng
// @Description Cập nhật thông tin khách hàng
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "ID khách hàng"
// @Param customer body model.Customer true "Dữ liệu cập nhật" example({"full_name":"Nguyễn Văn D"})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /customers/{id} [put]
func (h *CustomerHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Update(uint(id), &customer); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Customer updated successfully", nil)
}

// Delete godoc
// @Summary Xóa khách hàng
// @Description Xóa một khách hàng
// @Tags customer
// @Produce json
// @Param id path int true "ID khách hàng"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /customers/{id} [delete]
func (h *CustomerHandler) Delete(c *gin.Context) {
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

	response.Success(c, "Customer deleted successfully", nil)
}
