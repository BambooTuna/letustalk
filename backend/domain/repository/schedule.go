package repository

import (
	"github.com/BambooTuna/letustalk/backend/domain"
	"time"
)

type ScheduleRepository interface {
	InsertSchedule(record domain.Schedule) error
	UpdateSchedule(record domain.Schedule) error

	//ResolveByScheduleId(scheduleId string) (*domain.Schedule, error)
	ResolveByParentAccountId(parentAccountId string, from time.Time, to time.Time) []*domain.Schedule
	ResolveByMyAccountId(myAccountId string, from time.Time, to time.Time) []*domain.Schedule
	//ResolveByChildAccountId(childAccountId string, from time.Time, to time.Time) []*domain.Schedule
}
