package models

import (
	"findmypal/config"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(username, password string) error {
	_, err := config.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}

func GetUser(username string) (string, error) {
	var hashedPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&hashedPassword)
	return hashedPassword, err
}
