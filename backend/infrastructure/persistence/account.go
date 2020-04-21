package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/jinzhu/gorm"
)

type AccountCredentialsRepositoryImpl struct {
	DBSession *gorm.DB
}

type AccountCredentialsRecord struct {
	AccountId string
	Mail      string
	Password  string
	Position  domain.AccountPosition
	Activated bool
}

func (AccountCredentialsRecord) TableName() string {
	return "account_credentials"
}

func (a AccountCredentialsRepositoryImpl) Insert(record *domain.AccountCredentials) error {
	accountCredentialsRecord := AccountCredentialsRecord{
		AccountId: record.AccountId,
		Mail:      record.Mail,
		Password:  record.Password,
		Position:  record.Position,
		Activated: record.Activated,
	}
	return a.DBSession.Create(accountCredentialsRecord).Error
}

func (a AccountCredentialsRepositoryImpl) UpdateActivated(accountId string, activated bool) error {
	if err := a.DBSession.Model(&AccountCredentialsRecord{AccountId: accountId}).Where(fmt.Sprintf("activated IS NOT %t", activated)).Update("activated", activated).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (a AccountCredentialsRepositoryImpl) ResolveByMail(mail string) (*domain.AccountCredentials, error) {
	var result AccountCredentialsRecord
	sql := fmt.Sprintf("select * from account_credentials where mail = '%s'", mail)
	if err := a.DBSession.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	} else {
		return &domain.AccountCredentials{
			AccountId: result.AccountId,
			Mail:      result.Mail,
			Password:  result.Password,
			Position:  result.Position,
			Activated: result.Activated,
		}, nil
	}
}

func (a AccountCredentialsRepositoryImpl) ResolveByAccountId(accountId string) (*domain.AccountCredentials, error) {
	var result AccountCredentialsRecord
	sql := fmt.Sprintf("select * from account_credentials where account_id = '%s'", accountId)
	if err := a.DBSession.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	} else {
		return &domain.AccountCredentials{
			AccountId: result.AccountId,
			Mail:      result.Mail,
			Password:  result.Password,
			Position:  result.Position,
			Activated: result.Activated,
		}, nil
	}
}

type AccountDetailRepositoryImpl struct {
	DBSession *gorm.DB
}

type AccountDetailRecord struct {
	AccountId    string
	Name         string
	Introduction string
}

func (AccountDetailRecord) TableName() string {
	return "account_detail"
}

func (a AccountDetailRepositoryImpl) All(q config.QuantityLimit) []*domain.AccountDetail {
	var result []*AccountDetailRecord
	sql := fmt.Sprintf("select * from account_detail ORDER BY account_id desc Limit %d,%d", q.Drop(), q.Limit)
	a.DBSession.Raw(sql).Scan(&result)
	r := make([]*domain.AccountDetail, len(result))
	for i, e := range result {
		accountDetail := domain.AccountDetail{
			AccountId:    e.AccountId,
			Name:         e.Name,
			Introduction: e.Introduction,
		}
		r[i] = &accountDetail
	}
	return r
}

func (a AccountDetailRepositoryImpl) AllMentor(q config.QuantityLimit) []*domain.AccountDetail {
	var result []*AccountDetailRecord
	sql := fmt.Sprintf("select account_detail.account_id,account_detail.name,account_detail.introduction from account_credentials,account_detail where account_credentials.account_id = account_detail.account_id and account_credentials.position = 'mentor' ORDER BY account_detail.account_id desc Limit %d,%d", q.Drop(), q.Limit)
	a.DBSession.Raw(sql).Scan(&result)
	r := make([]*domain.AccountDetail, len(result))
	for i, e := range result {
		accountDetail := domain.AccountDetail{
			AccountId:    e.AccountId,
			Name:         e.Name,
			Introduction: e.Introduction,
		}
		r[i] = &accountDetail
	}
	return r
}

func (a AccountDetailRepositoryImpl) ResolveByAccountId(accountId string) (*domain.AccountDetail, error) {
	var result AccountDetailRecord
	sql := fmt.Sprintf("select * from account_detail where account_id = '%s'", accountId)
	if err := a.DBSession.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	} else {
		return &domain.AccountDetail{
			AccountId:    result.AccountId,
			Name:         result.Name,
			Introduction: result.Introduction,
		}, nil
	}
}
