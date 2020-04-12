package application

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"github.com/payjp/payjp-go/v1"
)

type InvoiceDetailUseCase struct {
	InvoiceDetailRepository repository.InvoiceDetailRepository
	PaymentService          *payjp.Service
}

func (i InvoiceDetailUseCase) GetInvoiceDetail(invoiceId string) (*domain.InvoiceDetail, error) {
	if invoiceDetail, err := i.InvoiceDetailRepository.ResolveByInvoiceId(invoiceId); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceDetailUseCase) IssueAnInvoice(amount int) (*domain.InvoiceDetail, error) {
	if invoiceDetail, err := domain.GenerateInvoiceDetail(amount); err != nil {
		return nil, err
	} else if err := i.InvoiceDetailRepository.Insert(invoiceDetail); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceDetailUseCase) MakePayment(invoiceId, token string) (*domain.InvoiceDetail, error) {
	if invoiceDetail, err := i.InvoiceDetailRepository.ResolveByInvoiceId(invoiceId); err != nil {
		return nil, err
	} else if invoiceDetail.Paid == true {
		return nil, config.Error("支払いずみ")
	} else if charge, err := i.CreatePaymentCharge(invoiceId, token, invoiceDetail.Amount); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else if err := i.InvoiceDetailRepository.Update(invoiceDetail.ChangePaidState(charge.Captured)); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceDetailUseCase) CreatePaymentCharge(invoiceId, token string, amount int) (*payjp.ChargeResponse, error) {
	return i.PaymentService.Charge.Create(amount, payjp.Charge{
		Currency:    "jpy",
		CardToken:   token,
		Capture:     true,
		Description: "Book: 'The Art of Community'",
		Metadata: map[string]string{
			"InvoiceId": invoiceId,
		},
	})
}
