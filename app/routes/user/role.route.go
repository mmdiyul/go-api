package user

import (
	"github.com/gin-gonic/gin"
	"go-api/app/controllers/user"
)

func registerRoleRoutes(r *gin.RouterGroup) {
	roles := r.Group("roles")

	roles.GET("", user.GetAllRoles)
	roles.GET(":id", user.GetRoleByID)
	roles.POST("", user.AddRole)
	roles.PUT(":id", user.UpdateRole)
	roles.DELETE(":id", user.DeleteRole)
}
