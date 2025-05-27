package packing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPackSizes = []int{250, 500, 1000, 2000, 5000}

func TestPackingResult_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		items    int
		expected map[int]int
		err      error
	}{
		{
			name:     "When 1 item ordered",
			items:    1,
			expected: map[int]int{250: 1},
			err:      nil,
		},
		{
			name:     "When 250 items ordered - matches exactly one of the existing packages",
			items:    250,
			expected: map[int]int{250: 1},
			err:      nil,
		},
		{
			name:     "When 251 items ordered",
			items:    251,
			expected: map[int]int{500: 1},
			err:      nil,
		},
		{
			name:     "When 501 items ordered",
			items:    501,
			expected: map[int]int{250: 1, 500: 1},
			err:      nil,
		},
		{
			name:     "When 12001 items ordered",
			items:    12001,
			expected: map[int]int{250: 1, 2000: 1, 5000: 2},
			err:      nil,
		},
		{
			name:     "Wrong - Invalid zero items",
			items:    0,
			expected: nil,
			err:      ErrInvalidItemCount,
		},
		{
			name:     "Wrong - Negative items",
			items:    -10,
			expected: nil,
			err:      ErrInvalidItemCount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PackingResult{}
			err := p.Calculate(tt.items, testPackSizes)
			assert.ErrorIs(t, err, tt.err)
			if tt.err == nil {
				assert.Equal(t, tt.expected, p.PacksUsed)
			}
		})
	}
}
