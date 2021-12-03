package persistence

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresItemRepository struct {
	db *gorm.DB
}

func NewPostgresItemRepository(db *gorm.DB) domain.ItemRepository {
	return &postgresItemRepository{db}
}

func (repo *postgresItemRepository) FindByID(itemID uint) (*domain.Item, error) {
	var item domain.Item
	tx := repo.db.First(&item, "id = ?", itemID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &item, nil
}
