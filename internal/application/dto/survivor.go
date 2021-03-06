package dto

import (
	"fmt"
	"net/http"

	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type SurvivorRead struct {
	Id       uint             `json:"id"`
	Name     string           `json:"name"`
	Age      uint             `json:"age"`
	Gender   domain.Gender    `json:"gender"`
	Position *domain.Location `json:"position"`
}

type SurvivorWrite struct {
	Name      string              `json:"name" validate:"required,min=2"`
	Age       uint                `json:"age" validate:"required,gte=0"`
	Gender    domain.Gender       `json:"gender" validate:"required"`
	Position  *domain.Location    `json:"position" validate:"required"`
	Inventory []*SurvivorResource `json:"inventory" validate:"required,min=1,dive,required"`
}

type SurvivorResource struct {
	ItemID   uint `json:"item_id" validate:"required"`
	Quantity uint `json:"quantity" validate:"required,gte=0"`
}

type SurvivorLastPosition struct {
	SurvivorID uint
	Location   *domain.Location
}

func NewSurvivorLastPosition(survID uint, loc *domain.Location) *SurvivorLastPosition {
	return &SurvivorLastPosition{SurvivorID: survID, Location: loc}
}

func ConvertToSurvivorRead(surv *domain.Survivor) *SurvivorRead {
	if surv == nil {
		return nil
	}

	return &SurvivorRead{
		Id:       surv.ID,
		Name:     surv.Name,
		Age:      surv.Age,
		Gender:   surv.Gender,
		Position: surv.LastLocation,
	}
}

type survivorBuilder struct {
	itemRepo domain.ItemRepository
}

func NewSurvivorBuilder(itemRepo domain.ItemRepository) *survivorBuilder {
	return &survivorBuilder{itemRepo}
}

func (builder *survivorBuilder) BuildSurvivor(sw *SurvivorWrite) (*domain.Survivor, error) {
	calculator := domain.NewInventoryPriceCalculator()

	var resources []*domain.Resource
	for _, survRes := range sw.Inventory {
		item, err := builder.findItem(survRes)
		if err != nil {
			return nil, err
		}

		calculator.AddItem(item, int32(survRes.Quantity))

		resources = append(resources, &domain.Resource{
			Item:     item,
			Quantity: survRes.Quantity,
		})
	}

	if err := calculator.ValidatePrice(); err != nil {
		msg := core.NewErrorMessage(ItemNotFound, err.Error(), http.StatusBadRequest)
		return nil, msg
	}

	surv := &domain.Survivor{
		Name:         sw.Name,
		Age:          sw.Age,
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

func (builder *survivorBuilder) findItem(survRes *SurvivorResource) (*domain.Item, error) {
	item, err := builder.itemRepo.FindByID(survRes.ItemID)

	if err != nil {
		detail := fmt.Sprintf("could not find item with ID %d", survRes.ItemID)
		msg := core.NewErrorMessage(ItemNotFound, detail, http.StatusBadRequest)
		return nil, msg
	}

	return item, nil
}
