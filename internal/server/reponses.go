package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOK(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     msg,
		"data":    data,
	})
}

func ResponseBadRequest(context *gin.Context, msg string, data interface{}) {
	context.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   data,
		"msg":     msg,
	})
}
