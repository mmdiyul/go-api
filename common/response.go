package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, error any) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
		"error":      error,
	})
}

func errorValidationResponse(ctx *gin.Context, statusCode int, error error) {
	errorResult := make(map[string]string)
	for _, fieldError := range error.(validator.ValidationErrors) {
		errorResult[fieldError.Field()] = fieldError.Tag()
	}
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    "Bad Request",
		"error":      errorResult,
	})
}
