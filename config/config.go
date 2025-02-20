package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	password := os.Getenv("DB_PASS")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", password, "localhost", "5432", "findmypal")
	DB, err = sql.Open("postgres", connStr) //TODO
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection not established:", err)
	}
}
