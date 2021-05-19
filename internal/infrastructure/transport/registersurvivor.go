package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
)

type registerSurvivorCtrl struct {
	ucase *application.RegisterSurvivor
}

func NewRegisterSurvivorCtrl(ucase *application.RegisterSurvivor) *registerSurvivorCtrl {
	return &registerSurvivorCtrl{ucase}
}

func (ctrl *registerSurvivorCtrl) Execute(c *gin.Context) {
	var survivor dto.SurvivorWrite
	if err := c.ShouldBindJSON(&survivor); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(&survivor)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"survivor": response})
}
