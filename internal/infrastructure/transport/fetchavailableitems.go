package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type fetchAvailableItemsCtrl struct {
	ucase *application.FetchAvailableItems
}

func NewFetchAvailableItemsCtrl(ucase *application.FetchAvailableItems) *fetchAvailableItemsCtrl {
	return &fetchAvailableItemsCtrl{ucase}
}

func (ctrl *fetchAvailableItemsCtrl) Execute(c *gin.Context) {
	filter := domain.NewItemFilter("", "", 0)
	c.ShouldBind(filter)
	response, err := ctrl.ucase.Invoke(filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
