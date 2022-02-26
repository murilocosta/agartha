package domain

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	MaxInventoryPriceAllowed = 100
)

type Inventory struct {
	gorm.Model
	Disabled bool

	// Has many Resource
	Resources []*Resource

	// Belongs to Survivor
	OwnerID uint
	Owner   *Survivor `gorm:"foreignKey:OwnerID"`
}

type InventoryRepository interface {
	SaveTransfer(from *Inventory, to *Inventory) error

	FindByOwnerID(ownerID uint) (*Inventory, error)
}

type InventoryPriceCalculator struct {
	totalPrice int32
}

func NewInventoryPriceCalculator() *InventoryPriceCalculator {
	return &InventoryPriceCalculator{totalPrice: 0}
}

func (c *InventoryPriceCalculator) AddItem(item *Item, quantity int32) {
	c.totalPrice = c.totalPrice + (item.Price * quantity)
}

func (c *InventoryPriceCalculator) ValidatePrice() error {
	if c.totalPrice > MaxInventoryPriceAllowed {
		msg := fmt.Sprintf(
			"inventory has exceeded the maximum size allowed: %d/%d",
			c.totalPrice,
			MaxInventoryPriceAllowed,
		)

		return errors.New(msg)
	}

	return nil
}

type InventoryService struct {
	repo InventoryRepository
}

func NewInventoryService(repo InventoryRepository) *InventoryService {
	return &InventoryService{repo}
}

func (s *InventoryService) TransferInventory(fromSurvID uint, toSurvID uint) error {
	fromInv, err := s.repo.FindByOwnerID(fromSurvID)
	if err != nil {
		return err
	}

	toInv, err := s.repo.FindByOwnerID(toSurvID)
	if err != nil {
		return err
	}

	toResMap := make(map[uint]*Resource)
	for _, toRes := range toInv.Resources {
		toResMap[toRes.ItemID] = toRes
	}

	for _, fromRes := range fromInv.Resources {
		if toRes, ok := toResMap[fromRes.ItemID]; ok {
			// Increase the quantity if the item already exists
			toRes.Quantity = toRes.Quantity + fromRes.Quantity
		} else {
			// If receiving inventory does not have the item, it should be created
			toInv.Resources = append(toInv.Resources, &Resource{
				InventoryID: toInv.ID,
				ItemID:      fromRes.ItemID,
				Quantity:    fromRes.Quantity,
			})
		}
	}

	if err := s.repo.SaveTransfer(fromInv, toInv); err != nil {
		return err
	}

	return nil
}
