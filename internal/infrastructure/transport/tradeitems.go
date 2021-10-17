package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
)

type tradeItemsCtrl struct {
	ucase *application.TradeItems
}

func NewTradeItemsCtrl(ucase *application.TradeItems) *tradeItemsCtrl {
	return &tradeItemsCtrl{ucase}
}

func (ctrl *tradeItemsCtrl) Execute(c *gin.Context) {
	var trade dto.TradeWrite
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(&trade)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
