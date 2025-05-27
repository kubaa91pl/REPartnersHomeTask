package shipment

import packing "home.excersise/internal/domain/packing"

type ShipmentRequest struct {
	Items int   `json:"items"`
	Packs []int `json:"packs"`
}

type ShipmentResult struct {
	PacksUsed map[int]int `json:"packs_used"`
}

func CreateShipment(input ShipmentRequest) ShipmentResult {
	p := &packing.PackingResult{}
	p.Calculate(input.Items, input.Packs)
	return ShipmentResult{
		PacksUsed: p.PacksUsed,
	}
}
