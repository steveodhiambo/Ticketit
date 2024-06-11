package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/steveodhiambo/ticket-it/cmd/api"
	"github.com/steveodhiambo/ticket-it/config"
	db "github.com/steveodhiambo/ticket-it/db"
	"log"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "ticket_it"
	password = "password123"
	dbname   = "ticket_it"
)

func main() {

	cfg := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		5432,
		config.Envs.DBName)

	// Intialize database
	db, err := db.NewPostgreSQLStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	//Initialize and conect to db
	initStorage(db)

	server := api.NewServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Start the database Connection
func initStorage(db *sql.DB) {
	//db.SetMaxOpenConns()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
}
