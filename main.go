// @title go-service-template API
// @version 0.1.0
// @description Minimal API for books
// @host localhost:8080
// @BasePath /
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	// BooksHandler serves a typed list of books.
	// @Summary List books
	// @Description Get a list of books
	// @Tags books
	// @Produce json
	// @Success 200 {array} Book
	// @Router /books [get]
	http.HandleFunc("/books", BooksHandler)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// BooksHandler returns a typed JSON array of Book
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := []Book{
		{ID: "1", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		{ID: "2", Title: "Lord of the Flies", Author: "William Golding"},
	}
	json.NewEncoder(w).Encode(books)
}
