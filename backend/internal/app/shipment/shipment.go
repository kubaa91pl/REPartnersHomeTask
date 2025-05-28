package shipment

import (
	"fmt"

	packing "home.excersise/internal/domain/packing"
	"home.excersise/internal/model"
	"home.excersise/internal/repository"
)

func CreateShipment(input model.ShipmentRequest, repo repository.ShipmentRepository) (model.ShipmentResult, error) {
	if len(input.Packs) == 0 {
		input.Packs = DefaultPackSizes
	}

	p := &packing.PackingResult{}
	err := p.Calculate(input.Items, input.Packs)
	if err != nil {
		return model.ShipmentResult{}, fmt.Errorf("error: CreateShipment fails due to: %w", err)
	}

	result := model.ShipmentResult{
		PacksUsed: p.PacksUsed,
	}

	_, err = repo.Save(&result)
	if err != nil {
		return model.ShipmentResult{}, fmt.Errorf("error: failed to save shipment result: %w", err)
	}

	return result, nil
}
