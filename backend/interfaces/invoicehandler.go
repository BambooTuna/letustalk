package interfaces

import (
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoiceDetailHandler struct {
	InvoiceDetailUseCase application.InvoiceDetailUseCase
}

func (i InvoiceDetailHandler) UnimplementedRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "UnimplementedRoute"})
	}
}
