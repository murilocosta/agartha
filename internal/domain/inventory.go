package domain

import "gorm.io/gorm"

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
