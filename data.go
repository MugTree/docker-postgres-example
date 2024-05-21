package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type User struct {
	Name     string
	Password string
}

func InitDb() error {

	var (
		dbname     = os.Getenv("POSTGRES_DB")
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

func GetUsers() ([]User, error) {
	users := []User{}
	rows, err := Db.Query(`SELECT user_name, password FROM app_user;`)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		u := User{}
		rows.Scan(&u.Name, &u.Password)
		users = append(users, u)
	}

	return users, nil

}

func AddUser(name string, password string) error {

	fmt.Println(name, password)

	_, err := Db.Exec("INSERT INTO app_user (user_name, password) VALUES($1, $2);", name, password)

	if err != nil {
		return err
	}

	return nil
}
