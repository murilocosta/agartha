package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type BindSurvivorID struct {
	SurvivorID uint `uri:"survivorId" binding:"required"`
}

type updateLastLocationCtrl struct {
	ucase *application.UpdateLastLocation
}

func NewUpdateLastLocationCtrl(ucase *application.UpdateLastLocation) *updateLastLocationCtrl {
	return &updateLastLocationCtrl{ucase}
}

func (ctrl *updateLastLocationCtrl) Execute(c *gin.Context) {
	var params BindSurvivorID
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, core.GetSystemError(err))
		return
	}

	var location domain.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(params.SurvivorID, &location)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"survivor": response})
}
