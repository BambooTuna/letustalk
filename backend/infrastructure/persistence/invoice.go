package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/domain"
	"gopkg.in/gorp.v1"
)

type InvoiceRepositoryImpl struct {
	DBSession *gorp.DbMap
}

type InvoiceRecord struct {
	InvoiceId string `db:"invoice_id"`
	Amount    int    `db:"amount"`
	Paid      bool   `db:"paid"`
}

func (i InvoiceRepositoryImpl) Insert(record *domain.Invoice) error {
	return i.DBSession.Insert(&InvoiceRecord{
		InvoiceId: record.InvoiceId,
		Amount:    record.Amount,
		Paid:      record.Paid,
	})
}

func (i InvoiceRepositoryImpl) Update(record *domain.Invoice) error {
	invoiceRecord := &InvoiceRecord{
		InvoiceId: record.InvoiceId,
		Amount:    record.Amount,
		Paid:      record.Paid,
	}
	if _, err := i.DBSession.Update(invoiceRecord); err != nil {
		return err
	} else {
		return nil
	}
}

func (i InvoiceRepositoryImpl) ResolveByInvoiceId(invoiceId string) (*domain.Invoice, error) {
	var result InvoiceRecord
	sql := fmt.Sprintf("select * from invoice_detail where invoice_id = '%s'", invoiceId)
	if err := i.DBSession.SelectOne(&result, sql); err != nil {
		return nil, err
	} else {
		return &domain.Invoice{
			InvoiceId: result.InvoiceId,
			Amount:    result.Amount,
			Paid:      result.Paid,
		}, nil
	}
}
