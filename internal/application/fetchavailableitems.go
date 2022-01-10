package application

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type FetchAvailableItems struct {
	itemRepo domain.ItemRepository
}

func NewFetchAvailableItems(itemRepo domain.ItemRepository) *FetchAvailableItems {
	return &FetchAvailableItems{itemRepo}
}

func (ucase *FetchAvailableItems) Invoke(filter *domain.ItemFilter) ([]*dto.ItemRead, error) {
	itemList, err := ucase.itemRepo.FindAll(filter)

	if err != nil {
		return nil, err
	}

	var resp []*dto.ItemRead
	for _, item := range itemList {
		respItem := dto.ConvertToItemRead(item)
		resp = append(resp, respItem)
	}
	return resp, nil
}
