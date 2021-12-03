package auth

import (
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type LoginSurvivor struct {
	credRepo domain.CredentialsRepository
}

func NewLoginSurvivor(credRepo domain.CredentialsRepository) *LoginSurvivor {
	return &LoginSurvivor{credRepo}
}

func (ucase *LoginSurvivor) Invoke(cred *AuthCredentials) (*SurvivorIdentity, bool) {
	user, err := ucase.credRepo.FindByUsername(cred.Username)
	if err != nil {
		return nil, false
	}

	if core.CheckPasswordHash(cred.Password, user.Password) {
		survIdentity := &SurvivorIdentity{
			SurvivorID: user.SurvivorID,
		}
		return survIdentity, true
	}

	return nil, false
}
