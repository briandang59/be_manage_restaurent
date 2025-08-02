package routes

import (
	"manage_restaurent/internal/handler"
	"github.com/gin-gonic/gin"
)

func FileRoutes(rg *gin.RouterGroup, h *handler.FileHandler) {
	rg.POST("/files/upload", h.UploadFile)
	rg.GET("/files", h.ListFiles)
	rg.GET("/files/:id", h.GetFile)
	rg.DELETE("/files/:id", h.DeleteFile)
} 