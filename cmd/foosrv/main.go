package main

import (
	"log"
	"net/http"

	"github.com/RabidFire/fooproject/pkg/api"
)

func main() {
	log.Println("Starting server...")

	http.HandleFunc("/", api.RootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
