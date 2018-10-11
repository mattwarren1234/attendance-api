package member

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"time"
)

func ByDate(d time.Time) error {
	return errors.New("not implemented")
}

type Member struct {
	ID    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Attendance struct {
	Member Member
	Events []*Event
}

// yay! yay!
type Event struct {
	Date      time.Time `json:"date"`
	Name      string    `json:"name"`
	Attendees *int      `json:"attendees, omitempty"` // optional field
}

// attendance count for each meeting date
func GetMemberCtByDay() ([]*Event, error) {
	db, err := getDB()
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

func attendanceCountByDay() ([]*Event, error) {
	db, err := getDB()
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

var _db *sql.DB

func getDB() (*sql.DB, error) {
	if _db == nil {
		var err error
		_db, err = sql.Open("postgres", "dbname=surj sslmode=disable")
		return _db, err
	} else {
		return _db, nil
	}
}

func getAttendanceByID(memberID int) (*Attendance, error) {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := db.Query("select members.member_id, first, last, email, phone, event.name, event.date from members left outer join attendance on (members.member_id = attendance.member_id) left outer join event on (attendance.event_id = event.event_id) where members.member_id=$1", memberID)
	if err != nil {
		return nil, err
	}
	attendance := new(Attendance)
	defer rows.Close()
	m := Member{}
	for rows.Next() {
		e := Event{}
		var eventName *string
		var eventDate *time.Time
		// because we are doing a let outer join, we might get null fields
		// something funky here where we overwrite the member info each time, but thats okay
		if err := rows.Scan(&m.ID, &m.First, &m.Last, &m.Email, &m.Phone, &eventName, &eventDate); err != nil {
			return nil, err
		}
		if eventName != nil && eventDate != nil {
			e.Name = *eventName
			e.Date = *eventDate
			attendance.Events = append(attendance.Events, &e)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	attendance.Member = m
	return attendance, nil
}

func GetAttendanceCountByDay(c echo.Context) error {
	events, err := attendanceCountByDay()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, events)
}

func GetAll(c echo.Context) error {
	members, err := getAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, members)
}

// func apiUser(c echo.Context) error {
// 	// response is ?
// 	response := []string{"some stuff"}
// 	val := map[string]interface{}{
// 		"user": response}
// 	return c.JSON(http.StatusOK, val)
// }

// returns 100 most recent members
func getAll() ([]*Member, error) {
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select member_id, first, last, email, phone from members order by member_id desc limit 100")
	if err != nil {
		return nil, err
	}
	var members []*Member
	defer rows.Close()
	for rows.Next() {
		m := Member{}
		if err := rows.Scan(&m.ID, &m.First, &m.Last, &m.Email, &m.Phone); err != nil {
			return nil, err
		}
		members = append(members, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return members, nil
}

func CreateEvent(c echo.Context) error {
	// given a list of user ids, an event name, and a date,
	// 1. insert name and date into events table, returning id
	// 2. insert into attendance the member ids and the event id
	db, err := getDB()
	if err != nil {
		return err
	}
	// {
	//  memberIDs : ["id1, id2"]
	//	date: someDate,
	//	name: someName,
	// }
	memberIDString := c.FormValue("memberIDs")
	date := c.FormValue("date")
	name := c.FormValue("name")
	if name == "" {
		return errors.New("missing name")
	}
	if date == "" {
		return errors.New("missing date")
	}
	if memberIDString == "" {
		return errors.New("missing memberIDString")
	}
	layout := "2014-09-12"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return err
	}
	var memberIDs []string
	err = json.Unmarshal([]byte(memberIDString), &memberIDs)
	if err != nil {
		return err
	}
	var eventID string
	err = db.QueryRow("insert into event(date, name) values($1, $2) returning event_id", parsedDate, name).Scan(&eventID)
	if err != nil {
		return err
	}
	formatted := strings.Join(memberIDs, ",")
	_, err = db.Exec("insert into attendance(event_id, member_d) values ($1, UNNEST(ARRAY[$2]))", eventID, formatted)
	return err
}

func GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	attendance, err := getAttendanceByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, attendance)
}
