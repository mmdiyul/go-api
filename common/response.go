package common

import "github.com/gin-gonic/gin"

func response(statusCode int, message string, data any) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	}
}

func Response(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, response(statusCode, message, data))
}
