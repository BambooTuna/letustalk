package domain

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/go-playground/validator"
)

type Invoice struct {
	InvoiceId string `json:"invoiceId" db:"invoice_id"`
	Amount    int    `json:"amount" db:"amount" validate:"gte=0"`
	Paid      bool   `json:"paid" db:"paid"`
}

func GenerateInvoice(amount int) (*Invoice, error) {
	uuid, err := config.GenerateUUID()
	if err != nil {
		return nil, err
	}
	details := Invoice{
		InvoiceId: uuid,
		Amount:    amount,
		Paid:      false,
	}
	if err := details.Validate(); err != nil {
		return nil, err
	}
	return &details, nil
}

func (i *Invoice) ChangePaidState(v bool) *Invoice {
	i.Paid = v
	return i
}

func (i *Invoice) Validate() error {
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
