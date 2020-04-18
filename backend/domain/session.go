package domain

import "encoding/json"

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
