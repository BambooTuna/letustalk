package persistence

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/jinzhu/gorm"
)

type ReservationRepositoryImpl struct {
	DBSession *gorm.DB
}

type ReservationRecord struct {
	ReservationId  string
	ChildAccountId string
	InvoiceId      string
}

func (ReservationRecord) TableName() string {
	return "reservation"
}

func (r ReservationRepositoryImpl) Insert(record *domain.Reservation) error {
	reservationRecord := ReservationRecord{
		ReservationId:  record.ReservationId,
		ChildAccountId: record.ChildAccountId,
		InvoiceId:      record.Invoice.InvoiceId,
	}
	return r.DBSession.Create(&reservationRecord).Error
}

//func (r ReservationRepositoryImpl) ResolveByReservationId(reservationId string) (*domain.Reservation, error) {
//
//}
//
//func (r ReservationRepositoryImpl) ResolveByChildAccountId(childAccountId string) []*domain.Reservation {
//
//}
