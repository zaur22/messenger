package db

import (
	"database/sql"
	"fmt"
	"log"
	"messenger/core/config"

	//driver for postres
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func init() {
	Connect()
}

func Connect() {
	var err error
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Print(err)
		}
	}
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%s host=%v port=%v",
		config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode, config.DBHost, config.DBPort)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
}

func MigrationUp() error {

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf(" with driver: %v", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func DropDB() error {

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"postgres", driver)
	if err != nil {
		return nil
	}
	return m.Drop()
}
