package dto

import "github.com/murilocosta/agartha/internal/domain"

type TradeWrite struct {
	Sender   *TradeInventoryWrite `json:"sender" validate:"required"`
	Receiver *TradeInventoryWrite `json:"receiver" validate:"required"`
}

type TradeInventoryWrite struct {
	SurvivorID uint                  `json:"survivor_id" validate:"required"`
	Items      []*TradeResourceWrite `json:"items" validate:"required,min=1,dive,required"`
}

type TradeResourceWrite struct {
	ResourceID uint `json:"resource_id" validate:"required"`
	Quantity   uint `json:"quantity" validate:"required,gte=0"`
}

type TradeRejectWrite struct {
	TradeID    uint   `uri:"tradeId" binding:"required"`
	Annotation string `json:"annotation"`
}

type TradeRead struct {
	TradeID uint               `json:"trade_id"`
	Status  domain.TradeStatus `json:"status"`
}

func ConvertToTradeRead(trade *domain.Trade) *TradeRead {
	if trade == nil {
		return nil
	}

	return &TradeRead{
		TradeID: trade.ID,
		Status:  trade.Status,
	}
}
