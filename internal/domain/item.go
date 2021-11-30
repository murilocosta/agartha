package domain

import "gorm.io/gorm"

type ItemRarity string

const (
	Common   ItemRarity = "Common"
	Uncommon ItemRarity = "Uncommon"
	Rare     ItemRarity = "Rare"
	Epic     ItemRarity = "Epic"
)

type Item struct {
	gorm.Model
	Name   string
	Icon   string
	Price  int32
	Rarity ItemRarity
}

type ItemRepository interface {
	FindByID(itemID uint) (*Item, error)
}
