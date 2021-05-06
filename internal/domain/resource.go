package domain

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Quantity uint

	// Has one Item
	ItemID uint
	Item   *Item

	// Belongs to Inventory
	InventoryID uint
}
