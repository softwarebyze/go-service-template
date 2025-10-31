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
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
