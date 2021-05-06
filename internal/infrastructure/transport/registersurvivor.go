package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
)

type registerSurvivorCtrl struct {
	ucase *application.RegisterSurvivor
}

func NewRegisterSurvivorCtrl(ucase *application.RegisterSurvivor) *registerSurvivorCtrl {
	return &registerSurvivorCtrl{ucase}
}

func (ctrl *registerSurvivorCtrl) Execute(c *gin.Context) {
	var survivor application.SurvivorWrite
	if err := c.ShouldBindJSON(&survivor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.ucase.Invoke(&survivor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"survivor": survivor})
}
