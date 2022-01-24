package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
)

type tradeItemsRejectCtrl struct {
	ucase *application.TradeItemsReject
}

func NewTradeItemsRejectCtrl(ucase *application.TradeItemsReject) *tradeItemsRejectCtrl {
	return &tradeItemsRejectCtrl{ucase}
}

func (ctrl *tradeItemsRejectCtrl) Execute(c *gin.Context) {
	var reject dto.TradeRejectWrite
	if err := c.ShouldBindUri(&reject); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
		return
	}

	if err := c.ShouldBindJSON(&reject); err != nil {
		c.JSON(http.StatusInternalServerError, core.GetSystemError(err))
	}

	response, err := ctrl.ucase.Invoke(&reject)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.GetErrorMessage(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
