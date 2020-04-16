package interfaces

import (
	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountDetailHandler struct {
	Session              session.Session
	AccountDetailUseCase application.AccountDetailUseCase
}

func (a AccountDetailHandler) GetAccountDetailsRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		page, e1 := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
		limit, e2 := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
		if e1 != nil || e2 != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "不正なクエリーパラメータです"})
			return
		}
		quantityLimit := config.QuantityLimit{Page: page, Limit: limit}
		ctx.JSON(http.StatusOK, a.AccountDetailUseCase.GetAccountDetails(quantityLimit))
	}
}

func (a AccountDetailHandler) GetMentorAccountDetailsRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		page, e1 := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
		limit, e2 := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
		if e1 != nil || e2 != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "不正なクエリーパラメータです"})
			return
		}
		quantityLimit := config.QuantityLimit{Page: page, Limit: limit}
		ctx.JSON(http.StatusOK, a.AccountDetailUseCase.GetMentorAccountDetails(quantityLimit))
	}
}

func (a AccountDetailHandler) GetAccountDetailRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		accountId := ctx.Param(paramKey)
		if accountDetail, err := a.AccountDetailUseCase.GetAccountDetail(accountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, accountDetail)
		}
	}
}
