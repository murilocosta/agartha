package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
)

type fetchSurvivorProfileCtrl struct {
	ucase *application.FetchSurvivorDetails
}

func NewFetchSurvivorProfileCtrl(ucase *application.FetchSurvivorDetails) *fetchSurvivorProfileCtrl {
	return &fetchSurvivorProfileCtrl{ucase}
}

func (ctrl *fetchSurvivorProfileCtrl) Execute(c *gin.Context) {
	claims, err := GetSurvivorIdentity(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(claims.SurvivorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
