package repository

import "github.com/BambooTuna/letustalk/backend/domain"

type ReservationRepository interface {
	Insert(record *domain.Reservation) error
	//ResolveByReservationId(reservationId string) (*domain.Reservation, error)
	//ResolveByChildAccountId(childAccountId string) []*domain.Reservation
}
