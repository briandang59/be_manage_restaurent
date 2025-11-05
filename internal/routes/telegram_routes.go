package routes

import (
	"manage_restaurent/internal/handler"

	"github.com/gin-gonic/gin"
)

func TelegramRoutes(rg *gin.RouterGroup, h *handler.TelegramHandler) {
	rg.POST("/telegram/send", h.Send)
}
