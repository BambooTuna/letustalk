package interfaces

import (
	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ScheduleHandler struct {
	Session         session.Session
	ScheduleUseCase application.ScheduleUseCase
}

type FreeScheduleResponseJson struct {
	ScheduleId string    `json:"scheduleId"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	UnitPrice  int       `json:"unitPrice"`
}

// GetFreeSchedule godoc
// @Summary GetFreeSchedule
// @Description GetFreeSchedule
// @Param accountId path string true "accountId"
// @Param from query string false "from"
// @Param to query string false "to"
// @Success 200 {array} FreeScheduleResponseJson
// @Failure 400 {object} json.ErrorMessageJson
// @Router /accounts/{accountId}/schedules [get]
func (s ScheduleHandler) GetFreeScheduleRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		accountId := ctx.Param(paramKey)
		defaultFrom := time.Now() //.Format("2006/01/02 15:04:05")
		defaultTo := defaultFrom.AddDate(0, 1, 0)

		layout := "20060102150405"
		if from, err := time.Parse(layout, ctx.Query("from")); err == nil {
			defaultFrom = from
		}
		if to, err := time.Parse(layout, ctx.Query("to")); err == nil {
			defaultTo = to
		}

		if defaultFrom.After(defaultTo) {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "from is after than to"})
		} else {
			result := s.ScheduleUseCase.GetFreeSchedule(accountId, defaultFrom, defaultTo)
			r := make([]*FreeScheduleResponseJson, len(result))
			for i, e := range result {
				r[i] = &FreeScheduleResponseJson{
					ScheduleId: e.ScheduleId,
					From:       e.From,
					To:         e.To,
					UnitPrice:  e.Detail.UnitPrice,
				}
			}
			ctx.JSON(http.StatusOK, r)
		}
	}
}

// Reserve godoc
// @Summary Reserve
// @Description Reserve
// @Param scheduleId path string true "scheduleId"
// @Param authorization header string true "authorization header"
// @Success 200
// @Failure 400 {object} json.ErrorMessageJson
// @Failure 403
// @Router /reservations/{scheduleId} [post]
func (s ScheduleHandler) ReserveRoute(paramKey string) func(ctx *gin.Context) {
	return s.Session.RequiredSession(func(ctx *gin.Context, token string) {
		scheduleId := ctx.Param(paramKey)
		accountSessionToken := domain.DecodeToAccountSessionToken(token)
		if err := s.ScheduleUseCase.Reserve(scheduleId, accountSessionToken.AccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.Status(http.StatusOK)
		}
	})
}
