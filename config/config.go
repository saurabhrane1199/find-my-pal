package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:pass@localhost:5432/findmypal?sslmode=disable") //TODO
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection not established:", err)
	}
}
