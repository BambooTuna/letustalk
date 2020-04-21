package repository

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"time"
)

type ScheduleRepository interface {
	InsertSchedule(record domain.Schedule) error
	CreateReservation(record domain.Schedule) error

	ResolveByScheduleId(scheduleId string) (*domain.Schedule, error)
	ResolveFreeScheduleByParentAccountId(parentAccountId string, from time.Time, to time.Time) []*domain.Schedule

	ResolveReservedScheduleByParentAccountId(parentAccountId string) []*domain.Schedule
	ResolveReservedScheduleByChildAccountId(childAccountId string) []*domain.Schedule
}
