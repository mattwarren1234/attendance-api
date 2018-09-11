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
