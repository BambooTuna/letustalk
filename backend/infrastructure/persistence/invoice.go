package persistence

import "gopkg.in/gorp.v1"

type InvoiceDetailRepositoryImpl struct {
	DBSession *gorp.DbMap
}
