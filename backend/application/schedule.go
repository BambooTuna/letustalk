package application

import (
	"errors"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"time"
)

type ScheduleUseCase struct {
	ScheduleRepository    repository.ScheduleRepository
	ReservationRepository repository.ReservationRepository
	InvoiceRepository     repository.InvoiceRepository
}

func (s ScheduleUseCase) GetFreeSchedule(accountId string, from time.Time, to time.Time) []*domain.Schedule {
	return s.ScheduleRepository.ResolveByParentAccountId(accountId, from, to)
}

func (s ScheduleUseCase) GetMySchedule(accountId string, from time.Time, to time.Time) []*domain.Schedule {
	return s.ScheduleRepository.ResolveByMyAccountId(accountId, from, to)
}

func (s ScheduleUseCase) Reserve(scheduleId, childAccountId string) error {
	if schedule, err := s.ScheduleRepository.ResolveByScheduleId(scheduleId); err != nil {
		println(1)
		return err
	} else if schedule.Reservation != nil {
		println(2)
		return errors.New("予約が埋まっています")
	} else if _, err := schedule.CreateReservation(childAccountId); err != nil {
		println(3)
		return err
	} else if err := s.InvoiceRepository.Insert(&schedule.Reservation.Invoice); err != nil {
		return err
	} else if err := s.ReservationRepository.Insert(schedule.Reservation); err != nil {
		return err
	} else if err := s.ScheduleRepository.UpdateSchedule(*schedule); err != nil {
		return err
	} else {
		return nil
	}
}
