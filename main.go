package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")

	http.HandleFunc("/", RootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
