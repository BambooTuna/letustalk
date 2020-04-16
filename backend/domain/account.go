package domain

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator"
)

type AccountDetail struct {
	AccountId    string `json:"accountId"`
	Name         string `json:"name"`
	Introduction string `json:"introduction" validate:"gte=0,lt=1000"`
}

type AccountCredentials struct {
	AccountId string
	Mail      string `validate:"required,email"`
	Password  string `validate:"gte=1,lt=255"`
}

func GenerateAccountCredentials(mail, plainPass string) (*AccountCredentials, error) {
	uuid, err := settings.GenerateUUID()
	if err != nil {
		return nil, err
	}
	accountCredentials := &AccountCredentials{
		AccountId: uuid,
		Mail:      mail,
		Password:  plainPass,
	}
	validate := validator.New()
	var errorMessages []config.CustomError
	if err := validate.Struct(accountCredentials); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, config.ValidateError(err.Field(), err.Tag()))
		}
		return nil, config.Errors(errorMessages)
	}

	encryptedPass, err := config.PasswordHash(plainPass)
	if err != nil {
		return nil, err
	}
	accountCredentials.Password = encryptedPass
	return accountCredentials, nil
}
