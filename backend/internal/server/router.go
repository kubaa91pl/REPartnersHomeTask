package server

import (
	"net/http"

	"home.excersise/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shipment", handler.CreateShipmentHandler)
	return mux
}
