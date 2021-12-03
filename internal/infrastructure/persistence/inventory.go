package persistence

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresInventoryRepository struct {
	db *gorm.DB
}

func NewPostgresInventoryRepository(db *gorm.DB) domain.InventoryRepository {
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

	tx := repo.db.Preload("Resources.Item").Preload(clause.Associations)
	rs := tx.Where("owner_id = ?", ownerID).First(&inv)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return &inv, nil
}
