package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOK(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"msg":     msg,
		"data":    data,
	})
}
