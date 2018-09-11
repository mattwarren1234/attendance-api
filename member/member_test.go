package member

import (
	"testing"
)

func TestGetMembers(t *testing.T) {
	members, err := GetMembers()
	if err != nil {
		t.Error(err)
	}
	if len(members) == 0 {
		t.Error("returned 0 members")
	}
}

var SampleUserID = 0

func TestGetAttendance(t *testing.T) {
	attendance, err := GetAttendance(SampleUserID)
	if err != nil {
		t.Error(err)
	}
	if len(attendance.Events) == 0 {
		t.Error("returned 0 members")
	}
}

func TestGetAttendanceCountByDay(t *testing.T) {
	events, err := GetAttendanceCountByDay()
	if err != nil {
		t.Error(err)
	}
	var hasAttendees bool
	for _, e := range events {
		if e.Attendees > 0 {
			hasAttendees = true
		}
	}
	if !hasAttendees {
		t.Error("no event has attendees")
	}
}

// func TestByDate(t *testing.T) {
// 	// given a date should give 0 or more users
// 	//give a date with no meeting should return nothing
// 	// given a date with people should return several things: f name l name, phone, and email.. AND number of times they've been there
// 	d := time.Now()
// 	err := ByDate(d)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// given a user (by id? sure) get the events they've been to
// given a date get a list of people that showed up
// given a date that was not a thing, an error if no event occurred that day

// what am i going to want first?
// for a person the number of events they've been to
// number of events they've attended
// list of events, with the date and the name.
// so i gues returned is a list with and event type. so i need an event type.
