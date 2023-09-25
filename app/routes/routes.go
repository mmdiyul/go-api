package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(r *gin.Engine, db *gorm.DB) {
	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "OK",
			"data":       nil,
		})
	})
}
