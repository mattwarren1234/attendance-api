package member

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestCreateEvent(t *testing.T) {
	v := url.Values{}
	v.Set("name", "testname")
	v.Set("date", "somedate")
	v.Set("memberIDs", "1,2,2")
	fmt.Println("encoded is ", v.Encode())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(v.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))
	e := echo.New()
	c := e.NewContext(req, rec)
	err := CreateEvent(c)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMemberCtByDay(t *testing.T) {
	eventCt, err := GetMemberCtByDay()
	if err != nil {
		t.Error(err)
	}
	if len(eventCt) == 0 {
		t.Error("0 events returned!")
	}

	// should be at least 1 day with valid ranges
	var hasValidCount bool
	for _, v := range eventCt {
		if v.Attendees != nil && *v.Attendees > 0 {
			hasValidCount = true
		}
	}
	if !hasValidCount {
		t.Error("doesn't have any events with a valid count")
	}
}

func TestGetAll(t *testing.T) {
	members, err := getAll()
	if err != nil {
		t.Error(err)
	}
	if len(members) == 0 {
		t.Error("returned 0 members")
	}
}

var SampleUserID = 0

func TestGetAttendanceByID(t *testing.T) {
	attendance, err := getAttendanceByID(668)
	if err != nil {
		t.Error(err)
	}
	if attendance.Member.First == "" {
		t.Error("member first name missing")
	}
	if len(attendance.Events) == 0 {
		t.Error("returned 0 members")
	}
}

func TestGetAttendanceCountByDay(t *testing.T) {
	events, err := attendanceCountByDay()
	if err != nil {
		t.Error(err)
	}
	var hasAttendees bool
	for _, e := range events {
		if e.Attendees != nil && *e.Attendees > 0 {
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
