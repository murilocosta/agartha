package application

import (
	"fmt"
	"net/http"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type FetchSurvivorTradeHistory struct {
	tradeRepo domain.TradeRepository
	itemRepo  domain.ItemRepository
}

func NewFetchSurvivorTradeHistory(tradeRepo domain.TradeRepository, itemRepo domain.ItemRepository) *FetchSurvivorTradeHistory {
	return &FetchSurvivorTradeHistory{tradeRepo, itemRepo}
}

func (ucase *FetchSurvivorTradeHistory) Invoke(survivorID uint) ([]*dto.TradeHistoryRead, error) {
	history, err := ucase.tradeRepo.FindTradeHistoryBySurvivor(survivorID)

	if err != nil {
		return nil, err
	}

	if len(history) == 0 {
		empty := make([]*dto.TradeHistoryRead, 0)
		return empty, nil
	}

	var resp []*dto.TradeHistoryRead
	for _, trade := range history {
		respItem, err := ucase.convertToTradeHistoryRead(trade)
		if err != nil {
			return nil, err
		}

		resp = append(resp, respItem)
	}
	return resp, nil
}

func (ucase *FetchSurvivorTradeHistory) convertToTradeHistoryRead(trade *domain.Trade) (*dto.TradeHistoryRead, error) {
	sender := &dto.TradeHistorySurvivorRead{
		SurvivorID: trade.Sender.Survivor.ID,
		Name:       trade.Sender.Survivor.Name,
	}

	receiver := &dto.TradeHistorySurvivorRead{
		SurvivorID: trade.Receiver.Survivor.ID,
		Name:       trade.Receiver.Survivor.Name,
	}

	senderItems, err := ucase.convertToTradeHistoryItemRead(trade.Sender.TradeResources)
	if err != nil {
		return nil, err
	}

	receiverItems, err := ucase.convertToTradeHistoryItemRead(trade.Sender.TradeResources)
	if err != nil {
		return nil, err
	}

	tradeHistory := &dto.TradeHistoryRead{
		TradeID:       trade.ID,
		Status:        trade.Status,
		Sender:        sender,
		Receiver:      receiver,
		SenderItems:   senderItems,
		ReceiverItems: receiverItems,
	}
	return tradeHistory, nil
}

func (ucase *FetchSurvivorTradeHistory) convertToTradeHistoryItemRead(resources []*domain.TradeResource) ([]*dto.TradeHistoryItemRead, error) {
	var senderItems []*dto.TradeHistoryItemRead
	for _, res := range resources {
		itemDetail, err := ucase.itemRepo.FindByID(res.ItemID)
		if err != nil {
			detail := fmt.Sprintf("could not find item with ID %d", res.ItemID)
			msg := core.NewErrorMessage(dto.ItemNotFound, detail, http.StatusInternalServerError)
			return nil, msg
		}

		senderItems = append(senderItems, &dto.TradeHistoryItemRead{
			ItemName:     itemDetail.Name,
			ItemQuantity: res.Quantity,
		})
	}

	return senderItems, nil
}
