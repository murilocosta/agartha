package application

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type TradeItems struct {
	survRepo  domain.SurvivorRepository
	tradeRepo domain.TradeRepository
}

func NewTradeItems(survRepo domain.SurvivorRepository, tradeRepo domain.TradeRepository) *TradeItems {
	return &TradeItems{survRepo, tradeRepo}
}

func (ucase *TradeItems) Invoke(tradeWrite *dto.TradeWrite) (*dto.TradeRead, error) {
	validate := validator.New()

	if err := validate.Struct(tradeWrite); err != nil {
		msg := core.NewErrorMessage(dto.RegisterSurvivorFailed, "trade registration has failed", http.StatusBadRequest)
		msg.AddErrorDetail(err, dto.ErrorDetailBuilder)
		return nil, msg
	}

	trade, err := ucase.buildTrade(tradeWrite)
	if err != nil {
		return nil, err
	}

	ucase.tradeRepo.Save(trade)

	return dto.ConvertToTradeRead(trade), nil
}

func (ucase *TradeItems) buildTrade(tw *dto.TradeWrite) (*domain.Trade, error) {
	senderInv, senderTradePoints, err := ucase.buildTradeInventory(tw.Sender)
	if err != nil {
		return nil, err
	}

	receiverInv, receiverTradePoints, err := ucase.buildTradeInventory(tw.Receiver)
	if err != nil {
		return nil, err
	}

	if senderTradePoints < receiverTradePoints {
		msg := core.NewErrorMessage(
			dto.TradeResourcePriceNotEquivalent,
			"the amount of items offered or requested are not equivalent",
			http.StatusBadRequest,
		)
		return nil, msg
	}

	trade := &domain.Trade{
		Status:   domain.TradeOpen,
		Sender:   senderInv,
		Receiver: receiverInv,
	}

	return trade, nil
}

func (ucase *TradeItems) buildTradeInventory(tInv *dto.TradeInventoryWrite) (*domain.TradeInventory, int32, error) {
	var resList []*domain.TradeResource

	tradePoints := int32(0)
	for _, item := range tInv.Items {
		res, err := ucase.tradeRepo.FindResourceByID(item.ResourceID)
		if err != nil {
			detail := fmt.Sprintf("could not find resource with ID %d", item.ResourceID)
			msg := core.NewErrorMessage(dto.TradeResourceNotFound, detail, http.StatusBadRequest)
			return nil, 0, msg
		}

		if res.Quantity < item.Quantity {
			detail := fmt.Sprintf("does not have enough '%s' to trade", res.Item.Name)
			msg := core.NewErrorMessage(dto.TradeResourceNotFound, detail, http.StatusBadRequest)
			return nil, 0, msg
		}
		tradePoints = tradePoints + (res.Item.Price * int32(item.Quantity))

		tRes := &domain.TradeResource{
			ItemID:   item.ResourceID,
			Item:     res,
			Quantity: item.Quantity,
		}
		resList = append(resList, tRes)
	}

	inventory := &domain.TradeInventory{
		SurvivorID:     tInv.SurvivorID,
		TradeResources: resList,
	}

	return inventory, tradePoints, nil
}
