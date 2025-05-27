package shipment

import (
	"fmt"

	packing "home.excersise/internal/domain/packing"
)

type ShipmentRequest struct {
	Items int   `json:"items"`
	Packs []int `json:"packs"`
}

type ShipmentResult struct {
	PacksUsed map[int]int `json:"packs_used"`
}

func CreateShipment(input ShipmentRequest) (ShipmentResult, error) {
	p := &packing.PackingResult{}
	err := p.Calculate(input.Items, input.Packs)
	if err != nil {
		return ShipmentResult{}, fmt.Errorf("error: CreateShipment fails due to: %w", err)
	}

	return ShipmentResult{PacksUsed: p.PacksUsed}, nil
}
