package infrastructure

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/infrastructure/transport"
)

type authMiddleware struct {
	realm               string
	jwtSecretKey        string
	tokenTimeout        time.Duration
	tokenRefreshTimeout time.Duration
}

func NewAuthMiddleware(realm string, jwtSecretKey string, tokenTimeout int32, tokenRefreshTimeout int32) *authMiddleware {
	return &authMiddleware{
		realm:               realm,
		jwtSecretKey:        jwtSecretKey,
		tokenTimeout:        time.Duration(tokenTimeout) * time.Second,
		tokenRefreshTimeout: time.Duration(tokenRefreshTimeout) * time.Second,
	}

}

func (mid *authMiddleware) Init(
	authenticator transport.AuthAuthenticateFunc,
	authorizator transport.AuthAuthorizeFunc,
) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           mid.realm,
		Key:             []byte(mid.jwtSecretKey),
		Timeout:         mid.tokenTimeout,
		MaxRefresh:      mid.tokenRefreshTimeout,
		IdentityKey:     auth.AuthIdentityKey,
		IdentityHandler: transport.CreateTokenIdentity,
		PayloadFunc:     transport.CreateTokenPayload,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		LoginResponse:   transport.FormatTokenResponse,
		RefreshResponse: transport.FormatTokenResponse,
		Unauthorized:    transport.FormatUnauthorizedResponse,
		TimeFunc:        time.Now,
	})

	if err != nil {
		return nil, err
	}

	if err := authMiddleware.MiddlewareInit(); err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
