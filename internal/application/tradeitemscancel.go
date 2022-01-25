package application

import (
	"net/http"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type TradeItemsCancel struct {
	tradeRepo domain.TradeRepository
}

func NewTradeItemsCancel(tradeRepo domain.TradeRepository) *TradeItemsCancel {
	return &TradeItemsCancel{tradeRepo}
}

func (ucase *TradeItemsCancel) Invoke(tradeCancel *dto.TradeCancelWrite) (*dto.TradeRead, error) {
	trade, err := ucase.tradeRepo.FindByID(tradeCancel.TradeID)
	if err != nil {
		return nil, err
	}

	if trade.Sender.SurvivorID != tradeCancel.SurvivorID {
		msg := core.NewErrorMessage(
			dto.TradeCanOnlyBeCancelledBySender,
			"only the sender can cancel a trade",
			http.StatusBadRequest,
		)
		return nil, msg
	}

	if trade.Status != domain.TradeOpen {
		msg := core.NewErrorMessage(
			dto.TradeStatusIsInvalid,
			"cannot cancel a trade that is not open",
			http.StatusBadRequest,
		)
		return nil, msg
	}

	err = ucase.tradeRepo.UpdateTradeStatusWithAnnotation(trade.ID, domain.TradeCancelled, tradeCancel.Annotation)
	if err != nil {
		return nil, err
	}

	return &dto.TradeRead{TradeID: trade.ID, Status: domain.TradeRejected}, nil
}
