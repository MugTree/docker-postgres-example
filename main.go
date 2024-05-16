package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	user, err := GetUser()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Name)
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env")
		return
	}

	err := InitDb()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Get("/", HandleGetUser)

}
