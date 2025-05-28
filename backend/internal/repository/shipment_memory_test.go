package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"home.excersise/internal/model"
)

func TestMemoryRepo_SaveAndGet(t *testing.T) {
	repo := NewMemoryRepository()

	t.Run("Happy Path - Save and Get successfully", func(t *testing.T) {
		result := &model.ShipmentResult{
			PacksUsed: map[int]int{250: 1, 500: 1},
		}
		id, err := repo.Save(result)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)

		fetched, err := repo.Get(id)
		assert.NoError(t, err)
		assert.Equal(t, result, fetched)
	})

	t.Run("Wrong - Save with nil result", func(t *testing.T) {
		_, err := repo.Save(nil)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrCannotSaveNilResult)
	})

	t.Run("Wrong - Get with empty ID", func(t *testing.T) {
		_, err := repo.Get("")
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrEmptyShipmentID)
	})

	t.Run("Wrong - Get non-existent ID", func(t *testing.T) {
		_, err := repo.Get("non-existent-id")
		assert.ErrorIs(t, err, ErrShipmentNotFound)
	})
}
