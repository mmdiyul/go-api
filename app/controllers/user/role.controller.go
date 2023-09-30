package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/bindings/user"
	"go-api/common"
	"log/slog"
	"net/http"
)

func GetAllRoles(ctx *gin.Context) {
	common.SuccessResponse(ctx, http.StatusOK, "OK", nil)
}

func GetRoleByID(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.SuccessResponse(ctx, http.StatusOK, "OK", nil)
}

func AddRole(ctx *gin.Context) {
	valid := common.ValidateInput(ctx, &user.RoleInputBinding{})
	if valid {
		common.SuccessResponse(ctx, http.StatusOK, "OK", nil)
	}
}

func UpdateRole(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.SuccessResponse(ctx, http.StatusOK, "OK", nil)
}

func DeleteRole(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.SuccessResponse(ctx, http.StatusOK, "OK", nil)
}
