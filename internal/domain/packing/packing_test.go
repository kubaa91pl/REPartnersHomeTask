package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPackSizes = []int{250, 500, 1000, 2000, 5000}

func TestPackingResult_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected map[int]int
		err      error
	}{
		{
			name:     "1 item",
			amount:   1,
			expected: map[int]int{250: 1},
			err:      nil,
		},
		{
			name:     "Exact 250",
			amount:   250,
			expected: map[int]int{250: 1},
			err:      nil,
		},
		{
			name:     "251 items",
			amount:   251,
			expected: map[int]int{500: 1},
			err:      nil,
		},
		{
			name:     "501 items",
			amount:   501,
			expected: map[int]int{250: 1, 500: 1},
			err:      nil,
		},
		{
			name:     "12001 items",
			amount:   12001,
			expected: map[int]int{250: 1, 2000: 1, 5000: 2},
			err:      nil,
		},
		{
			name:     "Invalid zero items",
			amount:   0,
			expected: nil,
			err:      ErrInvalidItemCount,
		},
		{
			name:     "Negative items",
			amount:   -10,
			expected: nil,
			err:      ErrInvalidItemCount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PackingResult{}
			err := p.Calculate(tt.amount, testPackSizes)
			assert.ErrorIs(t, err, tt.err)
			if tt.err == nil {
				assert.Equal(t, tt.expected, p.PacksUsed)
			}
		})
	}
}
