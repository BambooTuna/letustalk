package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/domain"
	"gopkg.in/gorp.v1"
)

type InvoiceRepositoryImpl struct {
	DBSession *gorp.DbMap
}

func (i InvoiceRepositoryImpl) Insert(record *domain.Invoice) error {
	return i.DBSession.Insert(record)
}

func (i InvoiceRepositoryImpl) Update(record *domain.Invoice) error {
	if _, err := i.DBSession.Update(record); err != nil {
		return err
	} else {
		return nil
	}
}

func (i InvoiceRepositoryImpl) ResolveByInvoiceId(invoiceId string) (*domain.Invoice, error) {
	var result domain.Invoice
	sql := fmt.Sprintf("select * from invoice_detail where invoice_id = '%s'", invoiceId)
	if err := i.DBSession.SelectOne(&result, sql); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
