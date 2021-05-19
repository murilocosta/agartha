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

func (u *UpdateLastLocation) Invoke(survID uint, loc *domain.Location) (*dto.SurvivorRead, error) {
	surv, err := u.survRepo.FindByID(survID)
	if err != nil {
		detail := fmt.Sprintf("could not find survivor with ID %d", survID)
		msg := core.NewErrorMessage(dto.SurvivorNotFound, detail, http.StatusNotFound)
		return nil, msg
	}

	validate := validator.New()

	if err := validate.Struct(loc); err != nil {
		msg := core.NewErrorMessage(dto.UpdateLastLocationFailed, "update last location failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.SurvivorWriteErrorBuilder)
		return nil, msg
	}

	err = u.survRepo.UpdateLastLocation(surv, loc)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToSurvivorRead(surv), nil
}
