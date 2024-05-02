package main

import (
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/steveodhiambo/ticket-it/config"
	"github.com/steveodhiambo/ticket-it/db"
	"log"
	"os"
)

func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Get os arguments
	cmd := os.Args[(len(os.Args) - 1)]

	// Run migrations
	if cmd == "up" {
		if err := m.Up(); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalf("Failed to apply migrations: %v", err)
			} else {
				log.Println("No migrations to apply")
			}
		} else {
			log.Println("Migrations applied successfully")
		}
	}

	// Tear down migrations
	if cmd == "down" {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalf("Failed to reverse migrations: %v", err)
			} else {
				log.Println("No migrations to reverse")
			}
		} else {
			log.Println("Migrations reversed successfully")
		}
	}

}
