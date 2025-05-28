package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"home.excersise/internal/model"
)

var (
	ErrCannotSaveNilResult = errors.New("cannot save nil result")
	ErrShipmentNotFound    = errors.New("shipment not found")
	ErrEmptyShipmentID     = errors.New("shipment ID must not be empty")
)

type ShipmentRepository interface {
	Save(result *model.ShipmentResult) (string, error)
	Get(id string) (*model.ShipmentResult, error)
}

type memoryRepo struct {
	mu   sync.RWMutex
	data map[string]*model.ShipmentResult
}

func NewMemoryRepository() ShipmentRepository {
	return &memoryRepo{
		data: make(map[string]*model.ShipmentResult),
	}
}

// Save stores a ShipmentResult in memory and assigns it a new UUID.
//
// Returns the generated UUID string on success. If the input result is nil,
// it returns ErrCannotSaveNilResult.
func (r *memoryRepo) Save(result *model.ShipmentResult) (string, error) {
	if result == nil {
		return "", ErrCannotSaveNilResult
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	id := uuid.New().String()
	result.ID = id
	r.data[id] = result
	return id, nil
}

// Get retrieves a ShipmentResult from memory by its ID.
//
// If the ID is empty, it returns ErrEmptyShipmentID.
// If no result is found for the given ID, it returns ErrShipmentNotFound.
func (r *memoryRepo) Get(id string) (*model.ShipmentResult, error) {
	if id == "" {
		return nil, ErrEmptyShipmentID
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	result, ok := r.data[id]
	if !ok {
		return nil, ErrShipmentNotFound
	}
	return result, nil
}
