package repository

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
)

type AccountCredentialsRepository interface {
	Insert(record *domain.AccountCredentials) error
	UpdateActivated(accountId string, activated bool) error
	ResolveByMail(mail string) (*domain.AccountCredentials, error)
	ResolveByAccountId(accountId string) (*domain.AccountCredentials, error)
}

type AccountDetailRepository interface {
	//Insert(record *domain.AccountDetail) error
	//Update(record *domain.AccountDetail) error
	All(q config.QuantityLimit) []*domain.AccountDetail
	AllMentor(q config.QuantityLimit) []*domain.AccountDetail
	ResolveByAccountId(accountId string) (*domain.AccountDetail, error)
}
