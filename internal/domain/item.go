package domain

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string
	Icon  string
	Price int32
}

type ItemRepository interface {
	FindByID(itemID uint) (*Item, error)
}
