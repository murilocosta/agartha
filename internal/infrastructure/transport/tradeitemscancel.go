package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
)

type tradeItemsCancelCtrl struct {
	ucase *application.TradeItemsCancel
}

func NewTradeItemsCancelCtrl(ucase *application.TradeItemsCancel) *tradeItemsCancelCtrl {
	return &tradeItemsCancelCtrl{ucase}
}

func (ctrl *tradeItemsCancelCtrl) Execute(c *gin.Context) {
	var cancel dto.TradeCancelWrite
	if err := c.ShouldBindUri(&cancel); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	if err := c.ShouldBindJSON(&cancel); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
	}

	claims, err := GetSurvivorIdentity(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
	}

	cancel.SurvivorID = claims.SurvivorID
	response, err := ctrl.ucase.Invoke(&cancel)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
