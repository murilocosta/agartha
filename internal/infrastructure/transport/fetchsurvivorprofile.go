package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/core"
)

type fetchSurvivorProfileCtrl struct {
	ucase *application.FetchSurvivorDetails
}

func NewFetchSurvivorProfileCtrl(ucase *application.FetchSurvivorDetails) *fetchSurvivorProfileCtrl {
	return &fetchSurvivorProfileCtrl{ucase}
}

func (ctrl *fetchSurvivorProfileCtrl) Execute(c *gin.Context) {
	survivorKey, ok := c.Keys["survivorID"]
	if !ok {
		errMsg := errors.New("malformed authentication token")
		c.JSON(http.StatusInternalServerError, core.GetSystemError(errMsg))
		return
	}

	claims, ok := survivorKey.(*auth.SurvivorIdentity)
	if !ok {
		errMsg := errors.New("corrupted authentication token")
		c.JSON(http.StatusInternalServerError, core.GetSystemError(errMsg))
		return
	}

	response, err := ctrl.ucase.Invoke(claims.SurvivorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
