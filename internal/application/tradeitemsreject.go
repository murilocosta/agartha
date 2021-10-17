package application

import (
	"net/http"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type TradeItemsReject struct {
	tradeRepo domain.TradeRepository
}

func NewTradeItemsReject(tradeRepo domain.TradeRepository) *TradeItemsReject {
	return &TradeItemsReject{tradeRepo}
}

func (ucase *TradeItemsReject) Invoke(tradeReject *dto.TradeRejectWrite) (*dto.TradeRead, error) {
	trade, err := ucase.tradeRepo.FindByID(tradeReject.TradeID)
	if err != nil {
		return nil, err
	}

	if trade.Status != domain.TradeOpen {
		msg := core.NewErrorMessage(
			dto.TradeStatusIsInvalid,
			"cannot reject a trade that is not open",
			http.StatusBadRequest,
		)
		return nil, msg
	}

	err = ucase.tradeRepo.UpdateTradeStatusWithAnnotation(trade.ID, domain.TradeRejected, tradeReject.Annotation)
	if err != nil {
		return nil, err
	}

	return &dto.TradeRead{TradeID: trade.ID, Status: domain.TradeRejected}, nil
}
