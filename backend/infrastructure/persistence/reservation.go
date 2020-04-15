package persistence

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"gopkg.in/gorp.v1"
)

type ReservationRepositoryImpl struct {
	DBSession *gorp.DbMap
}

type ReservationRecord struct {
	ReservationId  string `db:"reservation_id"`
	ChildAccountId string `db:"child_account_id"`
	InvoiceId      string `db:"invoice_id"`
}

func (r ReservationRepositoryImpl) Insert(record *domain.Reservation) error {
	reservationRecord := ReservationRecord{
		ReservationId:  record.ReservationId,
		ChildAccountId: record.ChildAccountId,
		InvoiceId:      record.Invoice.InvoiceId,
	}
	return r.DBSession.Insert(reservationRecord)
}

//func (r ReservationRepositoryImpl) ResolveByReservationId(reservationId string) (*domain.Reservation, error) {
//
//}
//
//func (r ReservationRepositoryImpl) ResolveByChildAccountId(childAccountId string) []*domain.Reservation {
//
//}
