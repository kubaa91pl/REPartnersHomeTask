package server

import (
	"net/http"

	"home.excersise/internal/handler"
)

func NewRouterWithCORS() http.Handler {
	return withCORS(newRouter())
}

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shipment", handler.CreateShipmentHandler)
	return mux
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
