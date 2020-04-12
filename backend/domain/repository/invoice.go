package repository

import "github.com/BambooTuna/letustalk/backend/domain"

type InvoiceDetailRepository interface {
	Insert(record *domain.InvoiceDetail) error
	Update(record *domain.InvoiceDetail) error
	ResolveByInvoiceId(invoiceId string) (*domain.InvoiceDetail, error)
}
