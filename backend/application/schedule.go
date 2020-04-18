package application

import (
	"errors"
	"github.com/BambooTuna/letustalk/backend/config"
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"time"
)

type ScheduleUseCase struct {
	AccountCredentialsRepository repository.AccountCredentialsRepository
	ScheduleRepository           repository.ScheduleRepository
	ReservationRepository        repository.ReservationRepository
	InvoiceRepository            repository.InvoiceRepository
}

func (s ScheduleUseCase) CreateSchedule(parentAccountId string, position domain.AccountPosition, startTime time.Time, detail domain.ScheduleDetail) (*domain.Schedule, error) {
	if position != domain.Mentor {
		return nil, config.Error("権限がありません")
	} else if schedule, err := domain.GenerateSchedule(parentAccountId, startTime, detail); err != nil {
		return nil, err
	} else if err := s.ScheduleRepository.InsertSchedule(*schedule); err != nil {
		return nil, err
	} else {
		return schedule, nil
	}
}

func (s ScheduleUseCase) GetFreeSchedule(accountId string, from time.Time, to time.Time) []*domain.Schedule {
	return s.ScheduleRepository.ResolveFreeScheduleByParentAccountId(accountId, from, to)
}

func (s ScheduleUseCase) Reserve(scheduleId, childAccountId string) error {
	if accountCredentials, err := s.AccountCredentialsRepository.ResolveByAccountId(childAccountId); err != nil {
		return err
	} else if accountCredentials.Activated == false {
		return errors.New("アカウントがアクティベートされていません")
	} else if schedule, err := s.ScheduleRepository.ResolveByScheduleId(scheduleId); err != nil {
		return err
	} else if _, err := schedule.CreateReservation(childAccountId); err != nil {
		return err
	} else if err := s.ScheduleRepository.CreateReservation(*schedule); err != nil {
		return err
	} else {
		return nil
	}
}

func (s ScheduleUseCase) GetReservedReservationsByParentAccountId(accountId string) []*domain.Schedule {
	return s.ScheduleRepository.ResolveReservedScheduleByParentAccountId(accountId)
}

func (s ScheduleUseCase) GetReservedReservationsByChildAccountId(accountId string) []*domain.Schedule {
	return s.ScheduleRepository.ResolveReservedScheduleByChildAccountId(accountId)
}
