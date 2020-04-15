package persistence

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"gopkg.in/gorp.v1"
	"time"
)

type ScheduleRepositoryImpl struct {
	DBSession             *gorp.DbMap
	ReservationRepository repository.ReservationRepository
	InvoiceRepository     repository.InvoiceRepository
}

type ScheduleRecord struct {
	ScheduleId      string    `db:"schedule_id"`
	ParentAccountId string    `db:"parent_account_id"`
	From            time.Time `db:"from"`
	To              time.Time `db:"to"`
	ReservationId   string    `db:"reservation_id"`
}

func (s ScheduleRepositoryImpl) InsertSchedule(record domain.Schedule) error {
	var reservationId string
	if record.Reservation != nil {
		reservationId = record.Reservation.ReservationId
	}
	scheduleRecord := ScheduleRecord{
		ScheduleId:      record.ScheduleId,
		ParentAccountId: record.ParentAccountId,
		From:            record.From,
		To:              record.To,
		ReservationId:   reservationId,
	}
	return s.DBSession.Insert(scheduleRecord)
}

func (s ScheduleRepositoryImpl) UpdateSchedule(record domain.Schedule) error {
	var reservationId string
	if record.Reservation != nil {
		reservationId = record.Reservation.ReservationId
	}
	scheduleRecord := ScheduleRecord{
		ScheduleId:      record.ScheduleId,
		ParentAccountId: record.ParentAccountId,
		From:            record.From,
		To:              record.To,
		ReservationId:   reservationId,
	}
	if _, err := s.DBSession.Update(scheduleRecord); err != nil {
		return err
	} else {
		return nil
	}
}

//func (s ScheduleRepositoryImpl) ResolveByScheduleId(scheduleId string) (*domain.Schedule, error) {
//	return nil, nil
//}

func (s ScheduleRepositoryImpl) ResolveByParentAccountId(parentAccountId string, from time.Time, to time.Time) []*domain.Schedule {
	//type ResultRecord struct {
	//	ScheduleRecord
	//	ReservationRecord
	//	InvoiceRecord
	//}
	//
	//var result []ResultRecord
	//sql := fmt.Sprintf("select * from schedule where invoice_id = '%s'", parentAccountId)
	//s.DBSession.Select(&result, sql)
	//
	//r := make([]*domain.Schedule, len(result))
	//for i, e := range result {
	//	r[i] = &domain.Schedule{
	//		ScheduleId:      e.ScheduleId,
	//		ParentAccountId: e.ParentAccountId,
	//		From:            e.From,
	//		To:              e.To,
	//		Detail:          domain.ScheduleDetail{},
	//		Reservation:     &domain.Reservation{
	//			ReservationId:  e.ReservationRecord.ReservationId,
	//			ChildAccountId: "",
	//			Invoice:        nil,
	//		},
	//	}
	//}
	return nil
}

//func (s ScheduleRepositoryImpl) ResolveByChildAccountId(childAccountId string, from time.Time, to time.Time) []*domain.Schedule {
//	return nil
//}
