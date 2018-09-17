package member

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"time"
)

func ByDate(d time.Time) error {
	return errors.New("not implemented")
}

type Member struct {
	First string
	Last  string
	Email string
	Phone string
	Event string
}

type Attendance struct {
	Member Member
	Events []*Event
}

// yay! yay!
type Event struct {
	Date      time.Time
	Name      string
	Attendees int // optional field
}

// attendance count for each meeting date
func GetMemberCtByDay() ([]*Event, error) {
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	var events []*Event
	rows, err := db.Query("select name, date, count(attendance.member_id) from event join attendance on (attendance.event_id = event.event_id) group by name, date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := Event{}
		// something funky here where we overwrite the member info each time, but thats okay
		if err := rows.Scan(&e.Name, &e.Date, &e.Attendees); err != nil {
			return nil, err
		}
		events = append(events, &e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func GetAttendanceCountByDay() ([]*Event, error) {
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	var events []*Event
	rows, err := db.Query("select name, date, count(attendance.member_id) from event join attendance on (attendance.event_id = event.event_id) group by name, date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := Event{}
		// something funky here where we overwrite the member info each time, but thats okay
		if err := rows.Scan(&e.Name, &e.Date, &e.Attendees); err != nil {
			return nil, err
		}
		events = append(events, &e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func GetAttendance(memberID int) (*Attendance, error) {
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select first, last, email, phone, event.name, event.date from members join attendance on (members.member_id = attendance.member_id) join event on (attendance.event_id = event.event_id) where members.member_id=$1", memberID)
	if err != nil {
		return nil, err
	}
	attendance := new(Attendance)
	defer rows.Close()
	m := Member{}
	for rows.Next() {
		e := Event{}
		// something funky here where we overwrite the member info each time, but thats okay
		if err := rows.Scan(&m.First, &m.Last, &m.Email, &m.Phone, &e.Name, &e.Date); err != nil {
			return nil, err
		}
		attendance.Events = append(attendance.Events, &e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return attendance, nil
}

// returns 100 most recent members
func GetMembers() ([]*Member, error) {
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select first, last, email, phone from members order by member_id desc limit 100")
	if err != nil {
		return nil, err
	}
	var members []*Member
	defer rows.Close()
	for rows.Next() {
		m := Member{}
		if err := rows.Scan(&m.First, &m.Last, &m.Email, &m.Phone); err != nil {
			return nil, err
		}
		members = append(members, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return members, nil
}

func All() error {
	return errors.New("not implemented")
}
