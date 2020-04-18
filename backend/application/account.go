package application

import (
	"github.com/BambooTuna/go-server-lib/authentication"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
)

type AccountCredentialsUseCase struct {
	AccountCredentialsRepository repository.AccountCredentialsRepository
	ActivatorUseCase             authentication.ActivatorUseCase
}

func (a AccountCredentialsUseCase) SignUp(mail, password string) (*domain.AccountCredentials, error) {
	if accountCredentials, err := domain.NewAccountCredentials(mail, password); err != nil {
		return nil, err
	} else if err := a.AccountCredentialsRepository.Insert(accountCredentials); err != nil {
		return nil, config.Error("メール使用ずみ")
	} else if _, err := a.ActivatorUseCase.IssueCode(accountCredentials.AccountId, accountCredentials.Mail); err != nil {
		return accountCredentials, nil
	} else {
		return accountCredentials, nil
	}
}

func (a AccountCredentialsUseCase) IssueActivateCode(accountId string) error {
	if accountCredentials, err := a.AccountCredentialsRepository.ResolveByAccountId(accountId); err != nil {
		return err
	} else if accountCredentials.Activated == true {
		return config.Error("すでにアクティベートされています")
	} else if _, err := a.ActivatorUseCase.IssueCode(accountCredentials.AccountId, accountCredentials.Mail); err != nil {
		return err
	} else {
		return nil
	}
}

func (a AccountCredentialsUseCase) ActivateAccount(code string) error {
	if accountId, err := a.ActivatorUseCase.Activate(code); err != nil {
		return err
	} else if err := a.AccountCredentialsRepository.UpdateActivated(accountId, true); err != nil {
		return err
	} else {
		return nil
	}
}

func (a AccountCredentialsUseCase) SignIn(mail, password string) (*domain.AccountCredentials, error) {
	if accountCredentials, err := a.AccountCredentialsRepository.ResolveByMail(mail); err != nil || accountCredentials.Accessible(password) == false {
		return nil, config.Error("NotFound")
	} else {
		return accountCredentials, nil
	}
}

type AccountDetailUseCase struct {
	AccountDetailRepository repository.AccountDetailRepository
}

func (a AccountDetailUseCase) GetAccountDetails(q config.QuantityLimit) []*domain.AccountDetail {
	return a.AccountDetailRepository.All(q)
}

func (a AccountDetailUseCase) GetMentorAccountDetails(q config.QuantityLimit) []*domain.AccountDetail {
	return a.AccountDetailRepository.AllMentor(q)
}

func (a AccountDetailUseCase) GetAccountDetail(accountId string) (*domain.AccountDetail, error) {
	if accountDetail, err := a.AccountDetailRepository.ResolveByAccountId(accountId); err != nil {
		return nil, config.Error("NotFound")
	} else {
		return accountDetail, nil
	}
}
