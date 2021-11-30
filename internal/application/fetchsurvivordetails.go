package application

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type FetchSurvivorDetails struct {
	survRepo domain.SurvivorRepository
}

func NewFetchSurvivorDetails(survRepo domain.SurvivorRepository) *FetchSurvivorDetails {
	return &FetchSurvivorDetails{survRepo}
}

func (ucase *FetchSurvivorDetails) Invoke(survivorId uint) (*dto.SurvivorRead, error) {
	surv, err := ucase.survRepo.FindByID(survivorId)

	if err != nil {
		return nil, err
	}

	respItem := dto.ConvertToSurvivorRead(surv)
	return respItem, nil
}
