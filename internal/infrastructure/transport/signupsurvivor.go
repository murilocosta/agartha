package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/core"
)

type survivorSignUpCtrl struct {
	ucase *auth.SignUpSurvivor
}

func NewSurvivorSignUpCtrl(ucase *auth.SignUpSurvivor) *survivorSignUpCtrl {
	return &survivorSignUpCtrl{ucase}
}

func (ctrl *survivorSignUpCtrl) Execute(c *gin.Context) {
	var register auth.AuthSignUp
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"survivor": response})
}
