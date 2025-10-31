package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("running go-service-template")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("go-service-template, https://github.com/softwarebyze/go-service-template"))
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[
			{"id": "1", "title": "To Kill a Mockingbird", "author": "Harper Lee"},
			{"id": "2", "title": "Lord of the Flies", "author": "William Golding"}
		]`))
	})
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
