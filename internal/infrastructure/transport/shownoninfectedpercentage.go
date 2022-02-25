package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/reports"
	"github.com/murilocosta/agartha/internal/core"
)

type showNonInfectedPercentageCtrl struct {
	ucase *reports.ShowNonInfectedPercentage
}

func NewShowNonInfectedPercentageCtrl(ucase *reports.ShowNonInfectedPercentage) *showNonInfectedPercentageCtrl {
	return &showNonInfectedPercentageCtrl{ucase}
}

func (ctrl *showNonInfectedPercentageCtrl) Execute(c *gin.Context) {
	response, err := ctrl.ucase.Invoke()
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": response})
}
