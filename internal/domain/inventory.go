package domain

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model

	// Has many Resource
	Resources []*Resource

	// Belongs to Survivor
	OwnerID uint
	Owner   *Survivor `gorm:"foreignKey:OwnerID"`
}
