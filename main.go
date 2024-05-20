package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if _, isContainer := os.LookupEnv("IS_CONTAINER"); isContainer {
		fmt.Println("env variables provided in compose file")
		return nil
	}
	fmt.Println("loading .env vars from .env file")
	return godotenv.Load(".env")
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	user, err := GetUser()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", user.Name)

}

func main() {

	err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = InitDb()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Get("/", HandleGetUser)
	r.Get("/*", http.NotFound)

	if err := http.ListenAndServe(":8000", r); err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}

}
