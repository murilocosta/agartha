package application

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type FetchSurvivorList struct {
	survRepo domain.SurvivorRepository
}

func NewFetchSurvivorList(survRepo domain.SurvivorRepository) *FetchSurvivorList {
	return &FetchSurvivorList{survRepo}
}

func (ucase *FetchSurvivorList) Invoke(filter *domain.SurvivorFilter) ([]*dto.SurvivorRead, error) {
	survList, err := ucase.survRepo.FindAll(filter)

	if err != nil {
		return nil, err
	}

	var resp []*dto.SurvivorRead
	for _, surv := range survList {
		respItem := dto.ConvertToSurvivorRead(surv)
		resp = append(resp, respItem)
	}
	return resp, nil
}
