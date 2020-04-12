package domain

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/go-playground/validator"
)

type InvoiceDetail struct {
	InvoiceId string `db:"invoice_id"`
	Amount    int    `db:"amount" validate:"gte=0"`
	Paid      bool   `db:"paid"`
}

func GenerateInvoiceDetail(amount int) (*InvoiceDetail, error) {
	uuid, err := config.GenerateUUID()
	if err != nil {
		return nil, err
	}
	details := InvoiceDetail{
		InvoiceId: uuid,
		Amount:    amount,
		Paid:      false,
	}
	if err := details.Validate(); err != nil {
		return nil, err
	}
	return &details, nil
}

func (i *InvoiceDetail) ChangePaidState(v bool) *InvoiceDetail {
	i.Paid = v
	return i
}

func (i *InvoiceDetail) Validate() error {
	validate := validator.New()
	var errorMessages []config.CustomError
	if err := validate.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, config.ValidateError(err.Field(), err.Tag()))
		}
		return config.Errors(errorMessages)
	}
	return nil
}
