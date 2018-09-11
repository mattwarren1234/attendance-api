package member

import (
	"testing"
)

func TestGetMembers(t *testing.T) {
	err := GetMembers()
	if err != nil {
		t.Error(err)
	}
}

func GetFullList() {

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
