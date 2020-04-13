package domain

type AccountDetail struct {
	AccountId    string `db:"account_id"`
	Name         string `db:"name"`
	Introduction string `db:"introduction" validate:"gte=0"`
}
