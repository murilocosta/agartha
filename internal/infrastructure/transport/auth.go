package transport

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/core"
)

type AuthIdentityFunc func(*gin.Context) interface{}

type AuthCheckFunc func(c *gin.Context) (interface{}, error)

type AuthAllowFunc func(data interface{}, c *gin.Context) bool

type AuthMessageFormatter func(c *gin.Context, code int, message string)

func CreateIdentity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &auth.SurvivorIdentity{
		SurvivorID: claims[auth.AuthIdentityKey].(uint),
	}
}

func FormatUnauthorizedResponse(c *gin.Context, code int, message string) {
	c.JSON(code, core.GetAuthError(message, code))
}
