package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"home.excersise/internal/app/shipment"

	"github.com/stretchr/testify/assert"
)

func TestCreateShipmentHandler(t *testing.T) {
	tests := []struct {
		name           string
		request        shipment.ShipmentRequest
		expectedStatus int
		expectedBody   shipment.ShipmentResult
		shouldFail     bool
	}{
		{
			name: "Happy Path",
			request: shipment.ShipmentRequest{
				Items: 501,
				Packs: []int{250, 500, 1000},
			},
			expectedStatus: http.StatusOK,
			expectedBody: shipment.ShipmentResult{
				PacksUsed: map[int]int{250: 1, 500: 1},
			},
		},
		{
			name: "Wrong - Bad request due to zero items",
			request: shipment.ShipmentRequest{
				Items: 0,
				Packs: []int{250, 500},
			},
			expectedStatus: http.StatusBadRequest,
			shouldFail:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req := httptest.NewRequest(http.MethodPost, "/shipment", bytes.NewReader(body))
			rr := httptest.NewRecorder()

			CreateShipmentHandler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if !tt.shouldFail {
				var result shipment.ShipmentResult
				err := json.NewDecoder(rr.Body).Decode(&result)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, result)
			}
		})
	}
}
