package main

import (
	"database/sql"
	"fmt"
	"os"
)

var Db *sql.DB

type User struct {
	Name     string
	Password string
}

func InitDb() error {

	var (
		dbname     = os.Getenv("POSTGRES_DATABASE")
		dbuser     = os.Getenv("POSTGRES_USER")
		dbpassword = os.Getenv("POSTGRES_PASSWORD")
		dbhost     = os.Getenv("POSTGRES_HOST")
		dbport     = os.Getenv("POSTGRES_PORT")
		uri        = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbuser, dbname, dbpassword, dbhost, dbport)
	)

	db, err := sql.Open("postgres", uri)

	if err != nil {
		return err
	}

	Db = db

	return nil
}

func GetUser() (User, error) {
	u := User{}
	err := Db.QueryRow("SELECT name, password FROM users WHERE name = 'Matt").Scan(&u.Name, &u.Password)
	return u, err
}
