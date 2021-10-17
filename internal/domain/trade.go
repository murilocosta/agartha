package domain

import (
	"gorm.io/gorm"
)

type TradeStatus string

const (
	TradeOpen     TradeStatus = "Open"
	TradeAccepted TradeStatus = "Accepted"
	TradeRejected TradeStatus = "Rejected"
)

type Trade struct {
	gorm.Model
	Status TradeStatus

	// Has one Sender
	SenderID uint
	Sender   *TradeInventory `gorm:"foreignKey:SenderID"`

	// Has one Receiver
	ReceiverID uint
	Receiver   *TradeInventory `gorm:"foreignKey:ReceiverID"`

	Annotation string
}

type TradeInventory struct {
	gorm.Model

	// Belongs to Survivor
	SurvivorID uint
	Survivor   *Survivor `gorm:"foreignKey:SurvivorID"`

	// Has many TradeResource
	TradeResources []*TradeResource `gorm:"foreignKey:InventoryID"`
}

type TradeResource struct {
	gorm.Model
	Quantity uint

	// Has one Resource
	ItemID uint
	Item   *Resource `gorm:"foreignKey:ItemID"`

	// Belongs to TradeInventory
	InventoryID uint
}

type TradeRepository interface {
	Save(trade *Trade) error
	UpdateTradeStatus(tradeID uint, status TradeStatus) error
	UpdateTradeStatusWithAnnotation(tradeID uint, status TradeStatus, annotation string) error
	UpdateResourceItem(resID uint, quantity uint) error
	FindByID(tradeID uint) (*Trade, error)
	FindResourceByID(resID uint) (*Resource, error)
}
