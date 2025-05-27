package handler

import (
	"encoding/json"
	"net/http"

	"home.excersise/internal/app/shipment"
	"home.excersise/internal/model"
	"home.excersise/internal/repository"
)

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
