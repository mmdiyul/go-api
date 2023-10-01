package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateInput(ctx *gin.Context, body any) interface{} {
	if err := ctx.ShouldBindJSON(body); err != nil {
		errorValidationResponse(ctx, http.StatusBadRequest, err)
		return nil
	}
	return body
}
