package shipment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPackSizes = []int{250, 500, 1000, 2000, 5000}

func TestCreateShipment(t *testing.T) {
	tests := []struct {
		name     string
		request  ShipmentRequest
		expected map[int]int
		fails    bool
	}{
		{
			name: "Happy Path",
			request: ShipmentRequest{
				Items: 242,
				Packs: testPackSizes,
			},
			expected: map[int]int{250: 1},
		},
		{
			name: "Wrong - Calculating packing result fails",
			request: ShipmentRequest{
				Items: 0,
				Packs: testPackSizes,
			},
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CreateShipment(tt.request)
			if tt.fails {
				assert.Error(t, err)
				assert.ErrorContains(t, err, "CreateShipment fails")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result.PacksUsed)
			}
		})
	}
}
