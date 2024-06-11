package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func NewPostgreSQLStorage(cfg string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(time.Minute * 5) // Idle Time for db connection before being closed
	db.SetMaxOpenConns(10)                 // Maximum number of connections to the database

	return db, nil
}
