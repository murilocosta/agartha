package application

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type FetchSurvivorInventory struct {
	invRepo domain.InventoryRepository
}

func NewFetchSurvivorInventory(invRepo domain.InventoryRepository) *FetchSurvivorInventory {
	return &FetchSurvivorInventory{invRepo}
}

func (ucase *FetchSurvivorInventory) Invoke(survivorId uint) (*dto.InventoryRead, error) {
	inv, err := ucase.invRepo.FindByOwnerID(survivorId)

	if err != nil {
		return nil, err
	}

	respItem := dto.ConvertToInventoryRead(inv)
	return respItem, nil
}
