package application

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
)

type AccountDetailUseCase struct {
	AccountDetailRepository repository.AccountDetailRepository
}

func (a AccountDetailUseCase) GetAll(q config.QuantityLimit) []*domain.AccountDetail {
	return a.AccountDetailRepository.All(q)
}

func (a AccountDetailUseCase) GetAllMentor(q config.QuantityLimit) []*domain.AccountDetail {
	return a.AccountDetailRepository.AllMentor(q)
}

func (a AccountDetailUseCase) GetAccountDetail(accountId string) (*domain.AccountDetail, error) {
	if accountDetail, err := a.AccountDetailRepository.ResolveByAccountId(accountId); err != nil {
		return nil, config.Error("NotFound")
	} else {
		return accountDetail, nil
	}
}
