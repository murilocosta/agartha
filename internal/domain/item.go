package domain

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/core"
)

type ItemRarity string

const (
	Common   ItemRarity = "Common"
	Uncommon ItemRarity = "Uncommon"
	Rare     ItemRarity = "Rare"
	Epic     ItemRarity = "Epic"
)

type ItemFilter struct {
	Name      string                `form:"name"`
	Sort      core.DatabaseSortType `form:"sort"`
	Page      int                   `form:"page"`
	PageItems int                   `form:"page_items"`
}

func NewItemFilter(name string, sort string, page int) *ItemFilter {
	return &ItemFilter{
		Name:      name,
		Sort:      core.DatabaseSortType(sort),
		Page:      page,
		PageItems: 5,
	}
}

type Item struct {
	gorm.Model
	Name   string
	Icon   string
	Price  int32
	Rarity ItemRarity
}

type ItemRepository interface {
	FindByID(itemID uint) (*Item, error)
	FindAll(filter *ItemFilter) ([]*Item, error)
}
