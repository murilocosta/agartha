package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
)

type flagInfectedSurvivorCtrl struct {
	ucase *application.FlagInfectedSurvivor
}

func NewFlagInfectedSurvivorCtrl(ucase *application.FlagInfectedSurvivor) *flagInfectedSurvivorCtrl {
	return &flagInfectedSurvivorCtrl{ucase}
}

func (ctrl *flagInfectedSurvivorCtrl) Execute(c *gin.Context) {
	var infection dto.ReportedInfection
	if err := c.ShouldBindJSON(&infection); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	if err := ctrl.ucase.Invoke(&infection); err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
