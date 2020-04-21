package interfaces

import (
	"net/http"
	"strconv"

	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
)

type AccountCredentialsHandler struct {
	Session                   session.Session
	AccountCredentialsUseCase application.AccountCredentialsUseCase
}

type SignRequestJson struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

// SignUp godoc
// @Summary SignUp
// @Description SignUp
// @Accept  json
// @Produce  json
// @Param SignRequestJson body SignRequestJson true "Mail&Password"
// @Success 200
// @Header 200 {string} set-authorization "ログイン用セッショントークン"
// @Failure 400 {object} json.ErrorMessageJson
// @Router /auth/signup [post]
func (a AccountCredentialsHandler) SignUpRoute() func(ctx *gin.Context) {
	return a.Session.SetSession(func(ctx *gin.Context) *string {
		var tokenString *string

		var signRequestJson SignRequestJson
		if err := ctx.BindJSON(&signRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if accountCredentials, err := a.AccountCredentialsUseCase.SignUp(signRequestJson.Mail, signRequestJson.Pass); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			accountSessionToken := domain.AccountSessionToken{AccountId: accountCredentials.AccountId, Position: accountCredentials.Position}.ToString()
			tokenString = &accountSessionToken
			ctx.Status(http.StatusOK)
		}
		return tokenString
	})
}

// ActivateAccount godoc
// @Summary ActivateAccount
// @Description アカウント有効化
// @Param code path string true "アクティベート用コード"
// @Success 200
// @Failure 400 {object} json.ErrorMessageJson
// @Router /activate/account/{code} [get]
func (a AccountCredentialsHandler) ActivateAccountRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		code := ctx.Param(paramKey)
		if err := a.AccountCredentialsUseCase.ActivateAccount(code); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}

// SendActivateMail godoc
// @Summary SendActivateMail
// @Description SendActivateMail
// @Param authorization header string true "authorization header"
// @Success 200
// @Failure 400 {object} json.ErrorMessageJson
// @Failure 403
// @Router /activate/account [put]
func (a AccountCredentialsHandler) SendActivateMailRoute() func(ctx *gin.Context) {
	return a.Session.RequiredSession(func(ctx *gin.Context, token string) {
		accountSessionToken := domain.DecodeToAccountSessionToken(token)
		if err := a.AccountCredentialsUseCase.IssueActivateCode(accountSessionToken.AccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	})
}

// SignIn godoc
// @Summary SignIn
// @Description SignIn
// @Accept  json
// @Produce  json
// @Param SignRequestJson body SignRequestJson true "Mail&Password"
// @Success 200
// @Header 200 {string} set-authorization "ログイン用セッショントークン"
// @Failure 400 {object} json.ErrorMessageJson
// @Router /auth/signin [post]
func (a AccountCredentialsHandler) SignInRoute() func(ctx *gin.Context) {
	return a.Session.SetSession(func(ctx *gin.Context) *string {
		var tokenString *string

		var signRequestJson SignRequestJson
		if err := ctx.BindJSON(&signRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if accountCredentials, err := a.AccountCredentialsUseCase.SignIn(signRequestJson.Mail, signRequestJson.Pass); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			accountSessionToken := domain.AccountSessionToken{AccountId: accountCredentials.AccountId, Position: accountCredentials.Position}.ToString()
			tokenString = &accountSessionToken
			ctx.Status(http.StatusOK)
		}
		return tokenString
	})
}

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

type AccountDetailResponseJson struct {
	AccountId    string `json:"accountId"`
	Name         string `json:"name"`
	Introduction string `json:"introduction" validate:"gte=0,lt=1000"`
}

// GetMentorAccountDetails godoc
// @Summary GetMentorAccountDetails
// @Description メンター詳細一覧取得
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {array} AccountDetailResponseJson
// @Failure 400 {object} json.ErrorMessageJson
// @Router /mentor/ [get]
func (a AccountDetailHandler) GetMentorAccountDetailsRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		page, e1 := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
		limit, e2 := strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
		if e1 != nil || e2 != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "不正なクエリーパラメータです"})
			return
		}
		quantityLimit := config.QuantityLimit{Page: page, Limit: limit}
		result := a.AccountDetailUseCase.GetMentorAccountDetails(quantityLimit)
		r := make([]*AccountDetailResponseJson, len(result))
		for i, e := range result {
			r[i] = &AccountDetailResponseJson{
				AccountId:    e.AccountId,
				Name:         e.Name,
				Introduction: e.Introduction,
			}
		}
		ctx.JSON(http.StatusOK, r)
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
