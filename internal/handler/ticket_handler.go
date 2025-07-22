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