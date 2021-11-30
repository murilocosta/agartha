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

type TradeHistorySurvivorRead struct {
	SurvivorID uint   `json:"id"`
	Name       string `json:"name"`
}

type TradeHistoryItemRead struct {
	ItemName     string `json:"item_name"`
	ItemQuantity uint   `json:"item_quantity"`
}

type TradeHistoryRead struct {
	TradeID       uint                      `json:"id"`
	Status        domain.TradeStatus        `json:"status"`
	Sender        *TradeHistorySurvivorRead `json:"sender"`
	Receiver      *TradeHistorySurvivorRead `json:"receiver"`
	SenderItems   []*TradeHistoryItemRead   `json:"sender_items"`
	ReceiverItems []*TradeHistoryItemRead   `json:"receiver_items"`
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
