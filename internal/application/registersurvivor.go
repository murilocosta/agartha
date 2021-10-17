package application

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type RegisterSurvivor struct {
	survRepo domain.SurvivorRepository
	itemRepo domain.ItemRepository
}

func NewRegisterSurvivor(survRepo domain.SurvivorRepository, itemRepo domain.ItemRepository) *RegisterSurvivor {
	return &RegisterSurvivor{survRepo, itemRepo}
}

func (ucase *RegisterSurvivor) Invoke(survWrite *dto.SurvivorWrite) (*dto.SurvivorRead, error) {
	validate := validator.New()

	if err := validate.Struct(survWrite); err != nil {
		msg := core.NewErrorMessage(dto.RegisterSurvivorFailed, "register survivor failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.ErrorDetailBuilder)
		return nil, msg
	}

	surv, err := ucase.buildSurvivor(survWrite)
	if err != nil {
		return nil, err
	}

	ucase.survRepo.Save(surv)

	return dto.ConvertToSurvivorRead(surv), err
}

func (ucase *RegisterSurvivor) buildSurvivor(sw *dto.SurvivorWrite) (*domain.Survivor, error) {
	var resources []*domain.Resource
	for _, invItem := range sw.Inventory {
		resource, err := ucase.buildItem(invItem)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	surv := &domain.Survivor{
		Name:         sw.Name,
		Gender:       sw.Gender,
		LastLocation: sw.Position,
		Infected:     false,
		Deceased:     false,
		Inventory: &domain.Inventory{
			Resources: resources,
		},
	}

	return surv, nil
}

func (ucase *RegisterSurvivor) buildItem(survRes *dto.SurvivorResource) (*domain.Resource, error) {
	item, err := ucase.itemRepo.FindByID(survRes.ItemID)

	if err != nil {
		detail := fmt.Sprintf("could not find item with ID %d", survRes.ItemID)
		msg := core.NewErrorMessage(dto.ItemNotFound, detail, http.StatusBadRequest)
		return nil, msg
	}

	res := &domain.Resource{
		Item:     item,
		Quantity: survRes.Quantity,
	}

	return res, nil
}
