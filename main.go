package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("running go-service-template")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("go-service-template, https://github.com/softwarebyze/go-service-template"))
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		books := []Book{
			{"1", "To Kill a Mockingbird", "Harper Lee"},
			{"2", "Lord of the Flies", "William Golding"},
		}
		json.NewEncoder(w).Encode(books)

	})
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
