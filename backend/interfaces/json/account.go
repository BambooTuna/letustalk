package json

import (
	"github.com/BambooTuna/letustalk/backend/domain"
)

type AccountDetailResponseJson struct {
	AccountId    string `json:"accountId"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

func ConvertToAccountDetailResponseJson(a *domain.AccountDetail) *AccountDetailResponseJson {
	return &AccountDetailResponseJson{AccountId: a.AccountId, Name: a.Name, Introduction: a.Introduction}
}

func ConvertToAccountDetailsResponseJson(a []*domain.AccountDetail) []*AccountDetailResponseJson {
	r := make([]*AccountDetailResponseJson, len(a))
	for i, e := range a {
		r[i] = ConvertToAccountDetailResponseJson(e)
	}
	return r
}
