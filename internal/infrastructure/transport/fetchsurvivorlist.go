package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type fetchSurvivorListCtrl struct {
	ucase *application.FetchSurvivorList
}

func NewFetchSurvivorListCtrl(ucase *application.FetchSurvivorList) *fetchSurvivorListCtrl {
	return &fetchSurvivorListCtrl{ucase}
}

func (ctrl *fetchSurvivorListCtrl) Execute(c *gin.Context) {
	filter := domain.NewSurvivorFilter("", "", 0)
	c.ShouldBind(filter)
	response, err := ctrl.ucase.Invoke(filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
