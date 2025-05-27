package model

type ShipmentRequest struct {
	Items int   `json:"items"`
	Packs []int `json:"packs"`
}

type ShipmentResult struct {
	ID        string      `json:"id"`
	PacksUsed map[int]int `json:"packs_used"`
}
