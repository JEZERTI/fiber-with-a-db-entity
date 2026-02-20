package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := "host=localhost port=5432 user=postgres password=12345 dbname=fitness_club sslmode=disable"

	var err error
	DB, err = sql.Open("posgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB is not connected", err)
	}

	log.Println("DB connected")
}
