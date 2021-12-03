package persistence

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
)

type postgresSurvivorRepository struct {
	db *gorm.DB
}

func NewPostgresSurvivorRepository(db *gorm.DB) domain.SurvivorRepository {
	return &postgresSurvivorRepository{db}
}

func (repo *postgresSurvivorRepository) Save(surv *domain.Survivor) error {
	tx := repo.db.Save(surv)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresSurvivorRepository) UpdateLastLocation(surv *domain.Survivor, loc *domain.Location) error {
	updateFields := &domain.Survivor{LastLocation: loc}

	tx := repo.db.Model(surv).Updates(updateFields)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresSurvivorRepository) FindByID(survID uint) (*domain.Survivor, error) {
	var surv domain.Survivor

	tx := repo.db.First(&surv, survID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &surv, nil
}

func (repo *postgresSurvivorRepository) FindAll(filter *domain.SurvivorFilter) ([]*domain.Survivor, error) {
	var survList []*domain.Survivor

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

	err := tx.Find(&survList).Commit().Error
	if err != nil {
		return nil, err
	}

	return survList, nil
}
