package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipment_Create(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected map[int]int
	}{
		{
			name:     "1 item",
			amount:   1,
			expected: map[int]int{250: 1},
		},
		{
			name:     "Exact 250",
			amount:   250,
			expected: map[int]int{250: 1},
		},
		{
			name:     "251 items",
			amount:   251,
			expected: map[int]int{500: 1},
		},
		{
			name:     "501 items",
			amount:   501,
			expected: map[int]int{250: 1, 500: 1},
		},
		{
			name:     "12001 items",
			amount:   12001,
			expected: map[int]int{250: 1, 2000: 1, 5000: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shipment := &Shipment{}
			shipment.Create(tt.amount, DefaultPackSizes)
			assert.Equal(t, tt.expected, shipment.PacksUsed)
		})
	}
}
