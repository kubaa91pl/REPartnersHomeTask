package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"home.excersise/internal/model"
)

type ShipmentRepository interface {
	Save(result *model.ShipmentResult) (string, error)
	Get(id string) (*model.ShipmentResult, error)
}

var ErrShipmentNotFound = errors.New("shipment not found")

type memoryRepo struct {
	data map[string]*model.ShipmentResult
	mu   sync.RWMutex
}

func NewMemoryRepository() ShipmentRepository {
	return &memoryRepo{
		data: make(map[string]*model.ShipmentResult),
	}
}

func (r *memoryRepo) Save(result *model.ShipmentResult) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := uuid.New().String()
	result.ID = id
	r.data[id] = result
	return id, nil
}

func (r *memoryRepo) Get(id string) (*model.ShipmentResult, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result, ok := r.data[id]
	if !ok {
		return nil, ErrShipmentNotFound
	}
	return result, nil
}
