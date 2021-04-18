package domain

import "time"

type TradeStatus string

const (
	Open     TradeStatus = "Open"
	Accepted TradeStatus = "Accepted"
	Rejected TradeStatus = "Rejected"
)

type TradeResource struct {
	Resource *Resource
	Quantity int32
}

type TradeInventory struct {
	Owner *Survivor
	Items []*TradeResource
}

type Trade struct {
	Sender    *TradeInventory
	Receiver  *TradeInventory
	Status    TradeStatus
	CreatedAt time.Time
}
