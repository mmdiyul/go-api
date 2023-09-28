package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/common"
	"log/slog"
	"net/http"
)

func GetAllRoles(ctx *gin.Context) {
	common.Response(ctx, http.StatusOK, "OK", nil)
}

func GetRoleByID(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.Response(ctx, http.StatusOK, "OK", nil)
}

func AddRole(ctx *gin.Context) {
	common.Response(ctx, http.StatusOK, "OK", nil)
}

func UpdateRole(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.Response(ctx, http.StatusOK, "OK", nil)
}

func DeleteRole(ctx *gin.Context) {
	slog.Info(fmt.Sprintf("Role ID: %v", ctx.Param("id")))
	common.Response(ctx, http.StatusOK, "OK", nil)
}
