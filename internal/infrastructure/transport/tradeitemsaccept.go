package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
)

type BindTradeID struct {
	TradeID uint `uri:"tradeId" binding:"required"`
}

type tradeItemsAcceptCtrl struct {
	ucase *application.TradeItemsAccept
}

func NewTradeItemsAcceptCtrl(ucase *application.TradeItemsAccept) *tradeItemsAcceptCtrl {
	return &tradeItemsAcceptCtrl{ucase}
}

func (ctrl *tradeItemsAcceptCtrl) Execute(c *gin.Context) {
	var params BindTradeID
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, core.GetSystemError(err))
		return
	}

	response, err := ctrl.ucase.Invoke(params.TradeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
