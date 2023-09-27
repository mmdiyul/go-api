package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/common"
	"gorm.io/gorm"
	"net/http"
)

func Register(r *gin.Engine, db *gorm.DB) {
	r.GET("", func(context *gin.Context) {
		common.Response(context, http.StatusOK, "OK", nil)
	})
}
