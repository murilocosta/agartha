package application

import (
	"fmt"
	"net/http"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type InventoryOperation bool

const (
	InsertResource InventoryOperation = true
	RemoveResource InventoryOperation = false
)

type TradeItemsAccept struct {
	tradeRepo domain.TradeRepository
}

func NewTradeItemsAccept(tradeRepo domain.TradeRepository) *TradeItemsAccept {
	return &TradeItemsAccept{tradeRepo}
}

func (ucase *TradeItemsAccept) Invoke(tradeID uint) (*dto.TradeRead, error) {
	trade, err := ucase.tradeRepo.FindByID(tradeID)
	if err != nil {
		return nil, err
	}

	if trade.Status != domain.TradeOpen {
		msg := core.NewErrorMessage(
			dto.TradeStatusIsInvalid,
			"cannot accept a trade that is not open",
			http.StatusBadRequest,
		)
		return nil, msg
	}

	err = ucase.writeOffStock(trade)
	if err != nil {
		return nil, err
	}

	err = ucase.tradeRepo.UpdateTradeStatus(trade.ID, domain.TradeAccepted)
	if err != nil {
		return nil, err
	}

	return &dto.TradeRead{TradeID: trade.ID, Status: domain.TradeAccepted}, nil
}

func (ucase *TradeItemsAccept) writeOffStock(trade *domain.Trade) error {
	// Remove items from sender
	err := ucase.changeItemsFromInventory(RemoveResource, trade.Sender.TradeResources)
	if err != nil {
		return err
	}

	// Add items to receiver
	err = ucase.changeItemsFromInventory(InsertResource, trade.Receiver.TradeResources)
	if err != nil {
		return err
	}

	return nil
}

func (ucase *TradeItemsAccept) changeItemsFromInventory(operation InventoryOperation, tradeResources []*domain.TradeResource) error {
	for _, tr := range tradeResources {
		res, err := ucase.tradeRepo.FindResourceByID(tr.ItemID)
		if err != nil {
			return err
		}

		var newQuantity uint
		if operation == InsertResource {
			newQuantity = res.Quantity + tr.Quantity
		} else if res.Quantity < tr.Quantity {
			msg := fmt.Sprintf("does not have enough '%s' to trade", res.Item.Name)
			return core.NewErrorMessage(dto.TradeResourceQuantityNotEnough, msg, http.StatusBadRequest)
		} else {
			newQuantity = res.Quantity - tr.Quantity
		}

		if err := ucase.tradeRepo.UpdateResourceItem(res.ID, newQuantity); err != nil {
			return err
		}
	}

	return nil
}
