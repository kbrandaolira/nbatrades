package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("html"))
	http.Handle("/", fs)

	http.HandleFunc("/teams/", TeamsHandler)

	log.Println("Initializing Server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
