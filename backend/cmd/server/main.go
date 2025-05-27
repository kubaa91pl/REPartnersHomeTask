package main

import (
	"log"
	"net/http"

	"home.excersise/internal/server"
)

func main() {
	log.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", server.NewRouter())
	if err != nil {
		log.Fatal("Server errors due to: " + err.Error())
	}
}
