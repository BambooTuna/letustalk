package interfaces

import (
	"github.com/BambooTuna/go-server-lib/session"
	"github.com/BambooTuna/letustalk/backend/application"
	"github.com/BambooTuna/letustalk/backend/interfaces/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InvoiceHandler struct {
	Session        session.Session
	InvoiceUseCase application.InvoiceUseCase
}

// GetInvoice godoc
// @Summary GetInvoice
// @Description GetInvoice
// @Param invoiceId path string true "invoiceId"
// @Success 200 {object} domain.Invoice
// @Failure 400 {object} json.ErrorMessageJson
// @Router /invoices/{invoiceId} [get]
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

// IssueAnInvoice godoc
// @Summary IssueAnInvoice
// @Description IssueAnInvoice
// @Param IssueAnInvoiceRequestJson body IssueAnInvoiceRequestJson true "IssueAnInvoiceRequestJson"
// @Success 200 {object} domain.Invoice
// @Failure 400 {object} json.ErrorMessageJson
// @Router /invoices/ [post]
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

// MakePayment godoc
// @Summary MakePayment
// @Description MakePayment
// @Param invoiceId path string true "invoiceId"
// @Param MakePaymentRequestJson body MakePaymentRequestJson true "決済サービスより発行されたトークン"
// @Success 200 {object} domain.Invoice
// @Failure 400 {object} json.ErrorMessageJson
// @Router /invoices/{invoiceId} [post]
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
