package domain

import (
	"testing"
	"time"
)

func TestGenerateScheduleSuccess(t *testing.T) {
	scheduleDetail := ScheduleDetail{UnitPrice: 1000}

	if result, err := GenerateSchedule("test", time.Date(2020, 1, 1, 0, 29, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(2020, 1, 1, 0, 30, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date")
	} else if result.To != time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date")
	}

	if result, err := GenerateSchedule("test", time.Date(2020, 1, 31, 23, 30, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date")
	} else if result.To != time.Date(2020, 2, 1, 0, 30, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date")
	}

	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	if result, err := GenerateSchedule("test", time.Date(2020, 1, 31, 23, 59, 59, 0, location), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(2020, 2, 1, 0, 0, 0, 0, location).In(time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date (%s)", result.From)
	} else if result.To != time.Date(2020, 2, 1, 0, 30, 0, 0, location).In(time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date (%s)", result.To)
	}

}

func TestScheduleCreateReservationSuccess(t *testing.T) {
	scheduleDetail := ScheduleDetail{UnitPrice: 1000}

	childAccountId := "test_id"
	if schedule, err := GenerateSchedule("test", time.Date(2020, 1, 1, 0, 29, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if _, err := schedule.CreateReservation(childAccountId); err != nil {
		t.Fatalf("failed test (CreateReservation): %#v", err)
	} else if schedule.Reservation.ChildAccountId != childAccountId {
		t.Fatalf("failed test (CreateReservation): Reservation#ChildAccountId does not match")
	} else if schedule.Reservation.Invoice.Amount != scheduleDetail.UnitPrice {
		t.Fatalf("failed test (CreateReservation): Invoice#Amount not equal")
	}

}
