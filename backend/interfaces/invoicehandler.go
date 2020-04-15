package interfaces

import (
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceUseCase application.InvoiceUseCase
}

func (i InvoiceHandler) GetInvoiceRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param(paramKey)
		if invoiceDetail, err := i.InvoiceUseCase.GetInvoice(invoiceId); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}

type IssueAnInvoiceRequestJson struct {
	Amount int `json:"amount"`
}

func (i InvoiceHandler) IssueAnInvoiceRoute() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var issueAnInvoiceRequestJson IssueAnInvoiceRequestJson
		if err := ctx.BindJSON(&issueAnInvoiceRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if invoiceDetail, err := i.InvoiceUseCase.IssueAnInvoice(issueAnInvoiceRequestJson.Amount); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}

type MakePaymentRequestJson struct {
	Token string `json:"token"`
}

func (i InvoiceHandler) MakePaymentRoute(paramKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		invoiceId := ctx.Param(paramKey)
		var makePaymentRequestJson MakePaymentRequestJson
		if err := ctx.BindJSON(&makePaymentRequestJson); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else if invoiceDetail, err := i.InvoiceUseCase.MakePayment(invoiceId, makePaymentRequestJson.Token); err != nil {
			ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, invoiceDetail)
		}
	}
}
