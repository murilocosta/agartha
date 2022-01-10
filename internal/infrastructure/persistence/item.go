package persistence

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/core"
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

func (repo *postgresItemRepository) FindAll(filter *domain.ItemFilter) ([]*domain.Item, error) {
	var itemList []*domain.Item

	tx := repo.db.Begin()
	if filter.Name != "" {
		likeClause := fmt.Sprintf("%s%%", strings.ToLower(filter.Name))
		tx = tx.Where("lower(name) LIKE ?", likeClause)
	}

	if filter.Sort == core.Ascendent {
		tx = tx.Order("name asc")
	}

	if filter.Sort == core.Descendent {
		tx = tx.Order("name desc")
	}

	if filter.Page != 0 {
		offset := (filter.Page - 1) * filter.PageItems
		tx = tx.Limit(filter.PageItems).Offset(offset)
	}

	err := tx.Find(&itemList).Commit().Error
	if err != nil {
		return nil, err
	}

	return itemList, nil
}
