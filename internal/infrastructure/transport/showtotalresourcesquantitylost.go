package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/reports"
	"github.com/murilocosta/agartha/internal/core"
)

type showTotalResourcesQuantityLost struct {
	ucase *reports.ShowTotalResourceQuantityLost
}

func NewShowTotalResourcesQuantityLostCtrl(ucase *reports.ShowTotalResourceQuantityLost) *showTotalResourcesQuantityLost {
	return &showTotalResourcesQuantityLost{ucase}
}

func (ctrl *showTotalResourcesQuantityLost) Execute(c *gin.Context) {
	response, err := ctrl.ucase.Invoke()
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": response})
}
