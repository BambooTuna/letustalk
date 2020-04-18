package domain

import (
	"encoding/json"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/go-playground/validator"
)

type AccountCredentials struct {
	AccountId string
	Mail      string          `validate:"required,email"`
	Password  string          `validate:"gte=1,lt=255"`
	Position  AccountPosition `validate:"required"`
	Activated bool
}

func GenerateAccountCredentials(mail, plainPass string) (*AccountCredentials, error) {
	uuid, err := settings.GenerateUUID()
	if err != nil {
		return nil, err
	}
	accountCredentials := &AccountCredentials{
		AccountId: uuid,
		Mail:      mail,
		Position:  General,
		Activated: false,
	}
	return accountCredentials.ResetPassword(plainPass)
}

func (a *AccountCredentials) Validate() (*AccountCredentials, error) {
	validate := validator.New()
	var errorMessages []config.CustomError
	if err := validate.Struct(a); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, config.ValidateError(err.Field(), err.Tag()))
		}
		return nil, config.Errors(errorMessages)
	}
	return a, nil
}

func (a AccountCredentials) Accessible(plainPass string) bool {
	if encryptedPass, err := config.PasswordHash(plainPass); err != nil {
		return false
	} else {
		return a.Password == encryptedPass
	}
}

func (a *AccountCredentials) ResetPassword(newPassword string) (*AccountCredentials, error) {
	a.Password = newPassword
	if _, err := a.Validate(); err != nil {
		return nil, err
	} else if encryptedPass, err := config.PasswordHash(newPassword); err != nil {
		return nil, err
	} else {
		a.Password = encryptedPass
		return a, nil
	}
}

func (a *AccountCredentials) ChangePosition(newPosition AccountPosition) *AccountCredentials {
	a.Position = newPosition
	return a
}

func (a *AccountCredentials) Activate() *AccountCredentials {
	a.Activated = true
	return a
}

type AccountPosition string

const (
	General AccountPosition = "general"
	Mentor  AccountPosition = "mentor"
)

type AccountSessionToken struct {
	AccountId string          `json:"account_id"`
	Position  AccountPosition `json:"position"`
}

func (a AccountSessionToken) ToString() string {
	json, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(json)
}

func DecodeToAccountSessionToken(s string) *AccountSessionToken {
	var accountSessionToken *AccountSessionToken
	err := json.Unmarshal([]byte(s), &accountSessionToken)
	if err != nil {
		return nil
	}
	return accountSessionToken
}

type AccountDetail struct {
	AccountId    string `json:"accountId"`
	Name         string `json:"name"`
	Introduction string `json:"introduction" validate:"gte=0,lt=1000"`
}
