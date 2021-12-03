package transport

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/auth"
)

type survivorLoginCtrl struct {
	ucase *auth.LoginSurvivor
}

func NewSurvivorLoginCtrl(ucase *auth.LoginSurvivor) *survivorLoginCtrl {
	return &survivorLoginCtrl{ucase}
}

func (ctrl *survivorLoginCtrl) HandlerFunc(c *gin.Context) (interface{}, error) {
	var credentials auth.AuthCredentials
	if err := c.ShouldBind(&credentials); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	if response, ok := ctrl.ucase.Invoke(&credentials); ok {
		return response, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
