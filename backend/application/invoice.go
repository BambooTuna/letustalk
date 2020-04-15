package application

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"github.com/payjp/payjp-go/v1"
)

type InvoiceUseCase struct {
	InvoiceRepository repository.InvoiceRepository
	PaymentService    *payjp.Service
}

func (i InvoiceUseCase) GetInvoice(invoiceId string) (*domain.Invoice, error) {
	if invoiceDetail, err := i.InvoiceRepository.ResolveByInvoiceId(invoiceId); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceUseCase) IssueAnInvoice(amount int) (*domain.Invoice, error) {
	if invoiceDetail, err := domain.GenerateInvoice(amount); err != nil {
		return nil, err
	} else if err := i.InvoiceRepository.Insert(invoiceDetail); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceUseCase) MakePayment(invoiceId, token string) (*domain.Invoice, error) {
	if invoiceDetail, err := i.InvoiceRepository.ResolveByInvoiceId(invoiceId); err != nil {
		return nil, err
	} else if invoiceDetail.Paid == true {
		return nil, config.Error("支払いずみ")
	} else if charge, err := i.CreatePaymentCharge(invoiceId, token, invoiceDetail.Amount); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else if err := i.InvoiceRepository.Update(invoiceDetail.ChangePaidState(charge.Captured)); err != nil {
		return nil, config.Error(config.CustomError(err.Error()))
	} else {
		return invoiceDetail, nil
	}
}

func (i InvoiceUseCase) CreatePaymentCharge(invoiceId, token string, amount int) (*payjp.ChargeResponse, error) {
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
