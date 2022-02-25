package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application/reports"
	"github.com/murilocosta/agartha/internal/core"
)

type showAverageResourcesPerSurvivor struct {
	ucase *reports.ShowAverageResourcesPerSurvivor
}

func NewShowAverageResourcesPerSurvivorCtrl(ucase *reports.ShowAverageResourcesPerSurvivor) *showAverageResourcesPerSurvivor {
	return &showAverageResourcesPerSurvivor{ucase}
}

func (ctrl *showAverageResourcesPerSurvivor) Execute(c *gin.Context) {
	response, err := ctrl.ucase.Invoke()
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": response})
}
