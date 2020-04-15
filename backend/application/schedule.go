package application

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"github.com/BambooTuna/letustalk/backend/domain/repository"
	"time"
)

type ScheduleUseCase struct {
	ScheduleRepository repository.ScheduleRepository
}

func (s ScheduleUseCase) GetFreeSchedule(accountId string, from time.Time, to time.Time) []*domain.Schedule {
	return s.ScheduleRepository.ResolveByParentAccountId(accountId, from, to)
}

func (s ScheduleUseCase) GetMySchedule(accountId string, from time.Time, to time.Time) []*domain.Schedule {
	return s.ScheduleRepository.ResolveByMyAccountId(accountId, from, to)
}
