package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/jinzhu/gorm"
)

type InvoiceRepositoryImpl struct {
	DBSession *gorm.DB
}

type InvoiceRecord struct {
	InvoiceId string
	Amount    int
	Paid      bool
}

func (InvoiceRecord) TableName() string {
	return "invoice_detail"
}

func (i InvoiceRepositoryImpl) Insert(record *domain.Invoice) error {
	invoiceRecord := InvoiceRecord{
		InvoiceId: record.InvoiceId,
		Amount:    record.Amount,
		Paid:      record.Paid,
	}
	return i.DBSession.Create(invoiceRecord).Error
}

func (i InvoiceRepositoryImpl) Update(record *domain.Invoice) error {
	invoiceRecord := InvoiceRecord{
		InvoiceId: record.InvoiceId,
		Amount:    record.Amount,
		Paid:      record.Paid,
	}
	if err := i.DBSession.Save(&invoiceRecord).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (i InvoiceRepositoryImpl) ResolveByInvoiceId(invoiceId string) (*domain.Invoice, error) {
	var result InvoiceRecord
	sql := fmt.Sprintf("select * from invoice_detail where invoice_id = '%s'", invoiceId)
	if err := i.DBSession.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	} else {
		return &domain.Invoice{
			InvoiceId: result.InvoiceId,
			Amount:    result.Amount,
			Paid:      result.Paid,
		}, nil
	}
}
