package repository

import "github.com/BambooTuna/letustalk/backend/domain"

type InvoiceRepository interface {
	Insert(record *domain.Invoice) error
	Update(record *domain.Invoice) error
	ResolveByInvoiceId(invoiceId string) (*domain.Invoice, error)
}
