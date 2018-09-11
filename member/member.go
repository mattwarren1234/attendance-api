package member

import (
	"database/sql"
	"errors"
	"fmt"
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
	Date time.Time
	Name string
}

func GetAttendance(id int) (Attendance, error) {
	// get date whateverjoined name of the event
	// err = db.Query("select * from attendance JOIN  where id=? ")
	// to do : merge with whatever?
	// yeah cuz the one thats a big mere ghouth
	return Attendance{}, nil
}

func GetMembers() error {
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	m := Member{}
	err = db.QueryRow("select first, last, email, phone from members limit 1").Scan(&m.First, &m.Last, &m.Email, &m.Phone)
	// err = db.QueryRow("Select first, last, email, phone, event from members limit 1").Scan(&m.First, &m.Last, &m.Email, &m.Phone, &m.Event)
	if err != nil {
		return err
	}
	fmt.Println("M IS", m)
	return nil
}

func All() error {
	return errors.New("not implemented")
}
