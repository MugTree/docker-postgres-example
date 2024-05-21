package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if val, _ := os.LookupEnv("IS_CONTAINER"); val == "true" {
		return nil
	}
	fmt.Println("running dev .env")
	return godotenv.Load(".env.development")
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.New("page").ParseFiles("./html/page.html"))

	if r.Method == "POST" {

		fmt.Println("this is a post")
		r.ParseForm()

		err := AddUser(r.FormValue("Name"), r.FormValue("Password"))

		if err != nil {
			log.Fatal(err)
		}

	}

	users, err := GetUsers()

	for _, u := range users {
		fmt.Println("user:", u.Name)
	}

	if err != nil {
		log.Fatal(err)
	}

	vm := make(map[string]interface{})
	vm["Users"] = users

	temp.ExecuteTemplate(w, "page.html", vm)

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

	r.Get("/", HandleUsers)
	r.Post("/", HandleUsers)
	r.Get("/*", http.NotFound)

	if err := http.ListenAndServe(":8000", r); err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}

}
