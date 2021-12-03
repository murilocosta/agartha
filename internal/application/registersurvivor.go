package application

import (
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

	survBuilder := dto.NewSurvivorBuilder(ucase.itemRepo)
	surv, err := survBuilder.BuildSurvivor(survWrite)
	if err != nil {
		return nil, err
	}

	ucase.survRepo.Save(surv)

	return dto.ConvertToSurvivorRead(surv), err
}
