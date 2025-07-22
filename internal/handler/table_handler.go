package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"manage_restaurent/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TableHanlder struct {
	svc *service.TableService
}

func NewTableHandler(s *service.TableService) *TableHanlder {
	return &TableHanlder{svc: s}
}

// GetAll godoc
// @Summary Lấy danh sách bàn
// @Description Lấy danh sách bàn trong nhà hàng
// @Tags table
// @Produce json
// @Param page query int false "Trang"
// @Param page_size query int false "Số lượng mỗi trang"
// @Success 200 {object} response.Body{data=[]model.Table}
// @Router /tables [get]
func (h *TableHanlder) GetAll(c *gin.Context) {
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
// @Summary Tạo mới bàn
// @Description Tạo mới một bàn
// @Tags table
// @Accept json
// @Produce json
// @Param table body model.Table true "Dữ liệu bàn" example({"name":"Bàn 1","capacity":4})
// @Success 200 {object} model.Table
// @Failure 400 {object} response.ErrorResponse
// @Router /tables [post]
func (h *TableHanlder) Create(c *gin.Context) {
	var table model.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Create(&table); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, table, nil)
}

// Update godoc
// @Summary Cập nhật bàn
// @Description Cập nhật thông tin bàn
// @Tags table
// @Accept json
// @Produce json
// @Param id path int true "ID bàn"
// @Param updates body object true "Dữ liệu cập nhật" example({"capacity":6})
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /tables/{id} [patch]
func (h *TableHanlder) Update(c *gin.Context) {
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
	response.Success(c, "Availibility updated successfully", nil)
}

// Delete godoc
// @Summary Xóa bàn
// @Description Xóa một bàn
// @Tags table
// @Produce json
// @Param id path int true "ID bàn"
// @Success 200 {object} response.Body
// @Failure 400 {object} response.ErrorResponse
// @Router /tables/{id} [delete]
func (h *TableHanlder) Delete(c *gin.Context) {
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
	response.Success(c, "Availibility deleted successfully", nil)
}
