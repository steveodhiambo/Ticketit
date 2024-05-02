package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/steveodhiambo/ticket-it/cmd/api"
	"github.com/steveodhiambo/ticket-it/config"
	db "github.com/steveodhiambo/ticket-it/db"
)

func main() {

	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
		Timeout:              time.Minute * 5, // Maximun time to wait for database connection
	}

	// Intialize database
	db, err := db.NewMySQLStorage(cfg)

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
