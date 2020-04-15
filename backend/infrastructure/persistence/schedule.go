package persistence

import (
	"fmt"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/jinzhu/gorm"
	"time"
)

type ScheduleRepositoryImpl struct {
	DBSession *gorm.DB
}

type ScheduleRecord struct {
	ScheduleId      string
	ParentAccountId string
	From            time.Time
	To              time.Time
	ReservationId   string
}

type ScheduleDetailRecord struct {
	ScheduleId string
	UnitPrice  int
}

func (ScheduleRecord) TableName() string {
	return "schedule"
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
	return s.DBSession.Create(&scheduleRecord).Error
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
	if err := s.DBSession.Save(&scheduleRecord).Error; err != nil {
		return err
	} else {
		return nil
	}
}

//func (s ScheduleRepositoryImpl) ResolveByScheduleId(scheduleId string) (*domain.Schedule, error) {
//	return nil, nil
//}

func (s ScheduleRepositoryImpl) ResolveByParentAccountId(parentAccountId string, from time.Time, to time.Time) []*domain.Schedule {
	type ResultRecord struct {
		ScheduleRecord
		ScheduleDetailRecord
	}

	sql := fmt.Sprintf("select schedule.*, schedule_detail.* from schedule join schedule_detail on schedule.schedule_id = schedule_detail.schedule_id where schedule.reservation_id IS NULL and schedule.parent_account_id = '%s' and schedule.from >= '%s' and schedule.to <= '%s' ORDER BY schedule.from asc", parentAccountId, from, to)
	var result []*ResultRecord
	s.DBSession.Raw(sql).Scan(&result)
	r := make([]*domain.Schedule, len(result))
	for i, e := range result {
		schedule := domain.Schedule{
			ScheduleId:      e.ScheduleRecord.ScheduleId,
			ParentAccountId: e.ParentAccountId,
			From:            e.From,
			To:              e.To,
			Detail:          domain.ScheduleDetail{UnitPrice: e.ScheduleDetailRecord.UnitPrice},
		}
		r[i] = &schedule
	}
	return r
}

func (s ScheduleRepositoryImpl) ResolveByMyAccountId(myAccountId string, from time.Time, to time.Time) []*domain.Schedule {
	type ResultRecord struct {
		ScheduleRecord
		ReservationRecord
		InvoiceRecord
	}

	sql := fmt.Sprintf("select schedule.*, reservation.*, invoice_detail.* from schedule left outer join reservation on schedule.reservation_id = reservation.reservation_id left outer join invoice_detail on reservation.invoice_id = invoice_detail.invoice_id Where schedule.parent_account_id = '%s' and schedule.from >= '%s' and schedule.to <= '%s' ORDER BY schedule.from asc", myAccountId, from, to)
	var result []*ResultRecord
	s.DBSession.Raw(sql).Scan(&result)
	r := make([]*domain.Schedule, len(result))
	for i, e := range result {
		schedule := domain.Schedule{
			ScheduleId:      e.ScheduleId,
			ParentAccountId: e.ParentAccountId,
			From:            e.From,
			To:              e.To,
			Detail:          domain.ScheduleDetail{},
		}
		if e.ScheduleRecord.ReservationId != "" {
			schedule.Reservation = &domain.Reservation{
				ReservationId:  e.ReservationRecord.ReservationId,
				ChildAccountId: e.ReservationRecord.ChildAccountId,
				Invoice: domain.Invoice{
					InvoiceId: e.InvoiceRecord.InvoiceId,
					Amount:    e.InvoiceRecord.Amount,
					Paid:      e.InvoiceRecord.Paid,
				},
			}
		}
		r[i] = &schedule
	}
	return r
}

//func (s ScheduleRepositoryImpl) ResolveByChildAccountId(childAccountId string, from time.Time, to time.Time) []*domain.Schedule {
//	return nil
//}
