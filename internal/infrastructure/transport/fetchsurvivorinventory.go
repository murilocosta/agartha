package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
)

type fetchSurvivorInventoryCtrl struct {
	ucase *application.FetchSurvivorInventory
}

func NewFetchSurvivorInventoryCtrl(ucase *application.FetchSurvivorInventory) *fetchSurvivorInventoryCtrl {
	return &fetchSurvivorInventoryCtrl{ucase}
}

func (ctrl *fetchSurvivorInventoryCtrl) Execute(c *gin.Context) {
	var params BindSurvivorID
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(params.SurvivorID)

	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
