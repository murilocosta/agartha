package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
)

type fetchSurvivorDetailsCtrl struct {
	ucase *application.FetchSurvivorDetails
}

func NewFetchSurvivorDetailsCtrl(ucase *application.FetchSurvivorDetails) *fetchSurvivorDetailsCtrl {
	return &fetchSurvivorDetailsCtrl{ucase}
}

func (ctrl *fetchSurvivorDetailsCtrl) Execute(c *gin.Context) {
	var params BindSurvivorID
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(params.SurvivorID)

	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
