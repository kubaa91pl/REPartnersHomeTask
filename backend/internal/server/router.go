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

var allowedOrigins = map[string]bool{
	"http://localhost:9000":         true,
	"https://your-app.onrender.com": true,
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
