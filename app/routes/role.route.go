package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/app/controllers"
)

func registerRoleRoutes(r *gin.RouterGroup) {
	roles := r.Group("roles")

	roles.GET("", controllers.GetAllRoles)
	roles.GET(":id", controllers.GetRoleByID)
	roles.POST("", controllers.AddRole)
	roles.PUT(":id", controllers.UpdateRole)
	roles.DELETE(":id", controllers.DeleteRole)
}
