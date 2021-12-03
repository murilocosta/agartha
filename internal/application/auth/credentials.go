package auth

import (
	"github.com/murilocosta/agartha/internal/application/dto"
)

const (
	AuthIdentityKey string = "survivorID"
	AuthTokenType   string = "bearer"
)

type AuthSignUp struct {
	AuthCredentials
	Survivor *dto.SurvivorWrite `json:"survivor" validate:"required"`
}

type AuthCredentials struct {
	Username string `json:"username" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type SurvivorIdentity struct {
	SurvivorID uint
}
