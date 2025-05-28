package handler

import (
	"encoding/json"
	"net/http"

	"home.excersise/internal/app/shipment"
	"home.excersise/internal/model"
	"home.excersise/internal/repository"
)

// CreateShipmentHandler handles HTTP POST requests to create a shipment.
//
// It decodes a JSON body into a ShipmentRequest, validates the input,
// invokes the shipment creation logic, and responds with the ShipmentResult
// as a JSON response. Returns HTTP 400 for invalid input or processing errors.
func CreateShipmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.ShipmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := shipment.CreateShipment(req, repository.NewMemoryRepository())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
