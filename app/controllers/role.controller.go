package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/bindings"
	"go-api/app/services"
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
	data := common.ValidateInput(ctx, &bindings.RoleInputBinding{})
	if data != nil {
		common.SuccessResponse(ctx, http.StatusOK, "OK", services.AddRole(data))
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
