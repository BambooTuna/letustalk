package interfaces

import (
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoiceDetailHandler struct {
	InvoiceDetailUseCase application.InvoiceDetailUseCase
}

func (i InvoiceDetailHandler) GetInvoiceDetailRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param(paramKey)
		if invoiceDetail, err := i.InvoiceDetailUseCase.GetInvoiceDetail(invoiceId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}

type IssueAnInvoiceRequestJson struct {
	Amount int `json:"amount"`
}

func (i InvoiceDetailHandler) IssueAnInvoiceRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var issueAnInvoiceRequestJson IssueAnInvoiceRequestJson
		if err := ctx.BindJSON(&issueAnInvoiceRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if invoiceDetail, err := i.InvoiceDetailUseCase.IssueAnInvoice(issueAnInvoiceRequestJson.Amount); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}

type MakePaymentRequestJson struct {
	Token string `json:"token"`
}

func (i InvoiceDetailHandler) MakePaymentRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param(paramKey)
		var makePaymentRequestJson MakePaymentRequestJson
		if err := ctx.BindJSON(&makePaymentRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if invoiceDetail, err := i.InvoiceDetailUseCase.MakePayment(invoiceId, makePaymentRequestJson.Token); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}
