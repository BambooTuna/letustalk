package domain

import (
	"testing"
	"time"
)

func TestGenerateScheduleSuccess(t *testing.T) {
	scheduleDetail := ScheduleDetail{UnitPrice: 1000}

	if result, err := GenerateSchedule("test", time.Date(3020, 1, 1, 0, 29, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(3020, 1, 1, 0, 30, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date")
	} else if result.To != time.Date(3020, 1, 1, 1, 0, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date")
	}

	if result, err := GenerateSchedule("test", time.Date(3020, 1, 31, 23, 30, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(3020, 2, 1, 0, 0, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date")
	} else if result.To != time.Date(3020, 2, 1, 0, 30, 0, 0, time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date")
	}

	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	if result, err := GenerateSchedule("test", time.Date(3020, 1, 31, 23, 59, 59, 0, location), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if result.From != time.Date(3020, 2, 1, 0, 0, 0, 0, location).In(time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): from date (%s)", result.From)
	} else if result.To != time.Date(3020, 2, 1, 0, 30, 0, 0, location).In(time.UTC) {
		t.Fatalf("failed test (GenerateSchedule): to date (%s)", result.To)
	}

}

func TestScheduleCreateReservationSuccess(t *testing.T) {
	scheduleDetail := ScheduleDetail{UnitPrice: 50}
	childAccountId := "test_id"
	if schedule, err := GenerateSchedule("test", time.Date(3020, 1, 1, 0, 29, 59, 0, time.UTC), scheduleDetail); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if _, err := schedule.CreateReservation(childAccountId); err != nil {
		t.Fatalf("failed test (CreateReservation): %#v", err)
	} else if schedule.Reservation.ChildAccountId != childAccountId {
		t.Fatalf("failed test (CreateReservation): Reservation#ChildAccountId does not match")
	} else if schedule.Reservation.Invoice.Amount != scheduleDetail.UnitPrice {
		t.Fatalf("failed test (CreateReservation): Invoice#Amount not equal")
	}
}

func TestScheduleCreateReservationFailed(t *testing.T) {
	childAccountId := "test_id"
	if schedule, err := GenerateSchedule("test", time.Date(3020, 1, 1, 0, 29, 59, 0, time.UTC), ScheduleDetail{UnitPrice: -1000}); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if _, err := schedule.CreateReservation(childAccountId); err == nil {
		t.Fatalf("failed test (CreateReservation): マイナスの単価は設定できない")
	}

	if schedule, err := GenerateSchedule("test", time.Date(3020, 1, 1, 0, 29, 59, 0, time.UTC), ScheduleDetail{UnitPrice: 49}); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if _, err := schedule.CreateReservation(childAccountId); err == nil {
		t.Fatalf("failed test (CreateReservation): 50未満の単価は設定できない")
	}

	if schedule, err := GenerateSchedule("test", time.Now().AddDate(0, 0, -1), ScheduleDetail{UnitPrice: 49}); err != nil {
		t.Fatalf("failed test (GenerateSchedule): %#v", err)
	} else if _, err := schedule.CreateReservation(childAccountId); err == nil {
		t.Fatalf("failed test (CreateReservation): 予約可能時間を超えた予約はできない")
	}
}
