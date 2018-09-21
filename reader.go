package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/lib/pq"
)

func phoneNumber() string {
	num := func() int {
		return rand.Intn(10)
	}
	// lol
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", num(), num(), num(), num(), num(), num(), num(), num(), num(), num())
}

func main() {
	csvFile, _ := os.Open("names.csv")
	records, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", "dbname=surj sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range records {
		first := v[0]
		last := v[1]
		email := fmt.Sprintf("%s.%s@gmail.com", first, last)
		phone := phoneNumber()
		_, err := db.Exec("INSERT into members(first,last,email,phone) VALUES ($1, $2, $3, $4)", first, last, email, phone)
		if err != nil {
			log.Fatal(err)
		}
	}

}
