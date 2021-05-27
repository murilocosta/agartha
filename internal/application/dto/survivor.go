package dto

import (
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type SurvivorRead struct {
	Id       uint             `json:"id"`
	Name     string           `json:"name"`
	Gender   domain.Gender    `json:"gender"`
	Position *domain.Location `json:"position"`
}

type SurvivorWrite struct {
	Name      string              `json:"name" validate:"required"`
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

func SurvivorWriteErrorBuilder(field string, errorType string) *core.ErrorDetail {
	return &core.ErrorDetail{
		Name:   field,
		Reason: errorType,
	}
}

func ConvertToSurvivorRead(surv *domain.Survivor) *SurvivorRead {
	if surv == nil {
		return nil
	}

	return &SurvivorRead{
		Id:       surv.ID,
		Name:     surv.Name,
		Gender:   surv.Gender,
		Position: surv.LastLocation,
	}
}
