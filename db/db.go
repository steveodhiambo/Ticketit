package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxIdleTime(time.Minute * 5) // Idle Time for db connection before being closed
	db.SetMaxOpenConns(10)                 // Maximum number of connections to the database

	return db, nil
}
