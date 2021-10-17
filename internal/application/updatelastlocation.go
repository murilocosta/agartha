package application

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type UpdateLastLocation struct {
	survRepo domain.SurvivorRepository
}

func NewUpdateLastLocation(survRepo domain.SurvivorRepository) *UpdateLastLocation {
	return &UpdateLastLocation{survRepo}
}

func (u *UpdateLastLocation) Invoke(lastLoc *dto.SurvivorLastPosition) (*dto.SurvivorRead, error) {
	surv, err := u.survRepo.FindByID(lastLoc.SurvivorID)
	if err != nil {
		detail := fmt.Sprintf("could not find survivor with ID %d", lastLoc.SurvivorID)
		msg := core.NewErrorMessage(dto.SurvivorNotFound, detail, http.StatusNotFound)
		return nil, msg
	}

	validate := validator.New()

	if err := validate.Struct(lastLoc.Location); err != nil {
		msg := core.NewErrorMessage(dto.UpdateLastLocationFailed, "update last location failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.ErrorDetailBuilder)
		return nil, msg
	}

	err = u.survRepo.UpdateLastLocation(surv, lastLoc.Location)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToSurvivorRead(surv), nil
}
