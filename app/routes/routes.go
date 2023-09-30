package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/app/routes/user"
	"go-api/common"
	"gorm.io/gorm"
	"net/http"
)

func Register(r *gin.Engine, db *gorm.DB) {
	apiV1 := r.Group("v1")
	apiV1.GET("", func(context *gin.Context) {
		common.SuccessResponse(context, http.StatusOK, "API Works!", nil)
	})

	user.RegisterUserModuleRoutes(apiV1)
}
