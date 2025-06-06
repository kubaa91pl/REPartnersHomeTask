package shipment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"home.excersise/internal/model"
	"home.excersise/internal/repository"
)

var testPackSizes = []int{250, 500, 1000, 2000, 5000}

func TestCreateShipment(t *testing.T) {
	tests := []struct {
		name              string
		request           model.ShipmentRequest
		expectedPacksUsed map[int]int
		fails             bool
	}{
		{
			name: "Happy Path",
			request: model.ShipmentRequest{
				Items: 242,
				Packs: testPackSizes,
			},
			expectedPacksUsed: map[int]int{250: 1},
		},
		{
			name: "Wrong - Calculating packing result fails",
			request: model.ShipmentRequest{
				Items: 0,
				Packs: testPackSizes,
			},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CreateShipment(tt.request, repository.NewMemoryRepository())
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPacksUsed, result.PacksUsed)
				assert.NotEmpty(t, result.ID)
			}
		})
	}
}
