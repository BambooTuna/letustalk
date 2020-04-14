package domain

type AccountDetail struct {
	AccountId    string `json:"accountId" db:"account_id"`
	Name         string `json:"name" db:"name"`
	Introduction string `json:"introduction" db:"introduction" validate:"gte=0"`
}
