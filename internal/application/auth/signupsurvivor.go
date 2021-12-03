package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type SignUpSurvivor struct {
	credRepo domain.CredentialsRepository
	itemRepo domain.ItemRepository
}

func NewSignUpSurvivor(credRepo domain.CredentialsRepository, itemRepo domain.ItemRepository) *SignUpSurvivor {
	return &SignUpSurvivor{credRepo, itemRepo}
}

func (ucase *SignUpSurvivor) Invoke(credWrite *AuthSignUp) (*dto.SurvivorRead, error) {
	validate := validator.New()

	if err := validate.Struct(credWrite); err != nil {
		msg := core.NewErrorMessage(dto.RegisterSurvivorFailed, "survivor's sign up failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.ErrorDetailBuilder)
		return nil, msg
	}

	survBuilder := dto.NewSurvivorBuilder(ucase.itemRepo)
	surv, err := survBuilder.BuildSurvivor(credWrite.Survivor)
	if err != nil {
		return nil, err
	}

	hashedPass, err := core.HashPassword(credWrite.Password)
	if err != nil {
		return nil, err
	}

	cred := &domain.Credentials{
		Username: credWrite.Username,
		Password: hashedPass,
		Survivor: surv,
	}
	ucase.credRepo.Save(cred)

	return dto.ConvertToSurvivorRead(surv), err
}
