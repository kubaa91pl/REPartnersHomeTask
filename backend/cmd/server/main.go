package main

import (
	"log"
	"net/http"
	"os"

	"home.excersise/internal/server"
)

func main() {
	hostname, _ := os.Hostname()
	log.Printf("Server running at http://%s:8080", hostname)

	err := http.ListenAndServe(":8080", server.NewRouterWithCORS())
	if err != nil {
		log.Fatal("Server errors due to: " + err.Error())
	}
}
