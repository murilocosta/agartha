package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/murilocosta/agartha/internal/domain"
)

type SurvivorWrite struct {
	Name      string              `json:"name" validate:"required"`
	Gender    domain.Gender       `json:"gender" validate:"required"`
	Position  *domain.Location    `json:"position" validate:"required"`
	Inventory []*SurvivorResource `json:"inventory" validate:"required,dive,required"`
}

type SurvivorResource struct {
	ItemID   uint `json:"item_id" validate:"required"`
	Quantity uint `json:"quantity" validate:"required,gte=0"`
}

type RegisterSurvivor struct {
	survRepo domain.SurvivorRepository
	itemRepo domain.ItemRepository
}

func NewRegisterSurvivor(survRepo domain.SurvivorRepository, itemRepo domain.ItemRepository) *RegisterSurvivor {
	return &RegisterSurvivor{survRepo, itemRepo}
}

func (ucase *RegisterSurvivor) Invoke(survWrite *SurvivorWrite) error {
	validate := validator.New()

	if err := validate.Struct(survWrite); err != nil {
		return err
	}

	surv, err := ucase.buildSurvivor(survWrite)
	if err != nil {
		return err
	}

	ucase.survRepo.Save(surv)
	return nil
}

func (ucase *RegisterSurvivor) buildSurvivor(sw *SurvivorWrite) (*domain.Survivor, error) {
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

func (ucase *RegisterSurvivor) buildItem(survRes *SurvivorResource) (*domain.Resource, error) {
	item, err := ucase.itemRepo.FindByID(survRes.ItemID)
	if err != nil {
		return nil, err
	}
	res := &domain.Resource{
		Item:     item,
		Quantity: survRes.Quantity,
	}
	return res, nil
}
