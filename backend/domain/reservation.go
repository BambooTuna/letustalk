package domain

import "github.com/BambooTuna/letustalk/backend/config"

type Reservation struct {
	ReservationId  string
	ChildAccountId string
	Invoice        Invoice
}

func GenerateReservation(childAccountId string, amount int) (*Reservation, error) {
	uuid, err := config.GenerateUUID()
	if err != nil {
		return nil, err
	}
	if invoice, err := GenerateInvoice(amount); err != nil {
		return nil, err
	} else {
		reservation := Reservation{
			ReservationId:  uuid,
			ChildAccountId: childAccountId,
			Invoice:        *invoice,
		}
		return &reservation, nil
	}
}
