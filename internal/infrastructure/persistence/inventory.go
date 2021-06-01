package persistence

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresInventoryRepository struct {
	db *gorm.DB
}

func NewPostgresInventoryRepository(db *gorm.DB) *postgresInventoryRepository {
	return &postgresInventoryRepository{db}
}

func (repo *postgresInventoryRepository) SaveTransfer(from *domain.Inventory, to *domain.Inventory) error {
	tx := repo.db.Begin()

	// Must disable the source inventory
	tx.Model(from).Updates(&domain.Inventory{Disabled: true})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	// Save the new items on the target inventory
	tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(to)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	return tx.Commit().Error
}

func (repo *postgresInventoryRepository) FindByOwnerID(ownerID uint) (*domain.Inventory, error) {
	var inv domain.Inventory

	err := repo.db.Where("owner_id = ?", ownerID).First(&inv).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&inv).Association("Resources").Find(&inv.Resources)
	if err != nil {
		return nil, err
	}

	return &inv, nil
}
