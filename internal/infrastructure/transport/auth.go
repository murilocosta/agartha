package transport

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/core"
)

type AuthAuthenticateFunc func(c *gin.Context) (interface{}, error)

type AuthAuthorizeFunc func(data interface{}, c *gin.Context) bool

func CreateTokenIdentity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	// The algorithm somehow transforms the uint ID value to float64
	survivorID := claims[auth.AuthIdentityKey].(float64)
	return &auth.SurvivorIdentity{
		SurvivorID: uint(survivorID),
	}
}

func CreateTokenPayload(data interface{}) jwt.MapClaims {
	survivorData := data.(*auth.SurvivorIdentity)
	return jwt.MapClaims{
		auth.AuthIdentityKey: survivorData.SurvivorID,
	}
}

func FormatTokenResponse(c *gin.Context, status int, token string, expire time.Time) {
	response := &auth.AuthResponse{
		TokenType:   auth.AuthTokenType,
		AccessToken: token,
		ExpiresIn:   expire.Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, response)
}

func FormatUnauthorizedResponse(c *gin.Context, status int, message string) {
	c.JSON(status, core.GetAuthError(message, status))
}

func GetSurvivorIdentity(c *gin.Context) (*auth.SurvivorIdentity, error) {
	survivorKey, ok := c.Keys["survivorID"]
	if !ok {
		return nil, errors.New("malformed authentication token")
	}

	claims, ok := survivorKey.(*auth.SurvivorIdentity)
	if !ok {
		return nil, errors.New("corrupted authentication token")
	}

	return claims, nil
}
