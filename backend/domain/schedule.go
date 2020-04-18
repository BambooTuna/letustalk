package domain

import (
	"github.com/BambooTuna/letustalk/backend/config"
	"time"
)

type Schedule struct {
	ScheduleId      string
	ParentAccountId string
	From            time.Time
	To              time.Time
	Detail          ScheduleDetail
	Reservation     *Reservation
}

func GenerateSchedule(parentAccountId string, startTime time.Time, detail ScheduleDetail) (*Schedule, error) {
	uuid, err := config.GenerateUUID()
	if err != nil {
		return nil, err
	}
	start := startTime.In(time.UTC)
	var minute int
	if start.Minute() >= 30 {
		minute = 30
	} else {
		minute = 0
	}
	from := time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), minute, 0, 0, start.Location()).Add(30 * time.Minute)
	to := from.Add(30 * time.Minute)
	schedule := Schedule{
		ScheduleId:      uuid,
		ParentAccountId: parentAccountId,
		From:            from,
		To:              to,
		Detail:          detail,
		Reservation:     nil,
	}
	return &schedule, nil
}

func (s *Schedule) CreateReservation(childAccountId string) (*Schedule, error) {
	now := time.Now().UTC()
	if now.After(s.From.Add(time.Duration(30) * time.Minute)) {
		return nil, config.Error(config.ReservationTimeHasPassed)
	} else if s.Reservation != nil {
		return nil, config.Error(config.ReservationIsFull)
	} else if reservation, err := GenerateReservation(childAccountId, s.Detail.UnitPrice); err != nil {
		return nil, err
	} else {
		s.Reservation = reservation
		return s, nil
	}
}

type ScheduleDetail struct {
	UnitPrice int
}
