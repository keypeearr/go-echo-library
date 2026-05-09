package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kylerequez/go-echo-library/src/utils"
)

var DB *sql.DB

func Connect() error {
	log.Println("Connecting to the database...")

	uri, err := utils.GetEnv("DATABASE_URI")
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return err
	}

	DB = db
	log.Println("Successfully connected to the database!")
	return nil
}

func Disconnect() error {
	if err := DB.Close(); err != nil {
		return fmt.Errorf("There was an error when disconnecting the database: %s", err.Error())
	}

	return nil
}
