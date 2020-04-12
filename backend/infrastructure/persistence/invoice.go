package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/domain"
	"gopkg.in/gorp.v1"
)

type InvoiceDetailRepositoryImpl struct {
	DBSession *gorp.DbMap
}

func (i InvoiceDetailRepositoryImpl) Insert(record *domain.InvoiceDetail) error {
	return i.DBSession.Insert(record)
}

func (i InvoiceDetailRepositoryImpl) Update(record *domain.InvoiceDetail) error {
	if _, err := i.DBSession.Update(record); err != nil {
		return err
	} else {
		return nil
	}
}

func (i InvoiceDetailRepositoryImpl) ResolveByInvoiceId(invoiceId string) (*domain.InvoiceDetail, error) {
	var result domain.InvoiceDetail
	sql := fmt.Sprintf("select * from invoice_detail where invoice_id = '%s'", invoiceId)
	if err := i.DBSession.SelectOne(&result, sql); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
