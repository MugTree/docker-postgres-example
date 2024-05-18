package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	user, err := GetUser()

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(user.Name))

}

func main() {

	fmt.Println(os.Environ())

	err := InitDb()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Get("/", HandleGetUser)
	http.ListenAndServe(os.Getenv("APP_PORT"), r)

}
