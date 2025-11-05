package handler

import (
	"manage_restaurent/internal/model"
	"manage_restaurent/internal/response"
	"manage_restaurent/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TelegramHandler struct {
	svc *service.TelegramService
}

func NewTelegramHandler(s *service.TelegramService) *TelegramHandler {
	return &TelegramHandler{svc: s}
}

func (h *TelegramHandler) Send(c *gin.Context) {
	var req model.TelegramSendRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.svc.SendMessage(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, resp, nil)
}
