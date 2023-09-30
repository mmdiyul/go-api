package user

import "github.com/gin-gonic/gin"

func RegisterUserModuleRoutes(r *gin.RouterGroup) {
	registerRoleRoutes(r)
}
