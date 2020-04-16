package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/jinzhu/gorm"
)

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
	var result []*domain.AccountDetail
	sql := fmt.Sprintf("select * from account_detail ORDER BY account_id desc Limit %d,%d", q.Drop(), q.Limit)
	a.DBSession.Raw(sql).Scan(&result)
	return result
}

func (a AccountDetailRepositoryImpl) AllMentor(q config.QuantityLimit) []*domain.AccountDetail {
	var result []*domain.AccountDetail
	sql := fmt.Sprintf("select account_detail.account_id,account_detail.name,account_detail.introduction from account_credentials,account_detail where account_credentials.account_id = account_detail.account_id and account_credentials.position = 'mentor' ORDER BY account_detail.account_id desc Limit %d,%d", q.Drop(), q.Limit)
	a.DBSession.Raw(sql).Scan(&result)
	return result
}

func (a AccountDetailRepositoryImpl) ResolveByAccountId(accountId string) (*domain.AccountDetail, error) {
	var result domain.AccountDetail
	sql := fmt.Sprintf("select * from account_detail where account_id = '%s'", accountId)
	if err := a.DBSession.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
