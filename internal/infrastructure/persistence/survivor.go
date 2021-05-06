package persistence

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresSurvivorRepository struct {
	db *gorm.DB
}

func NewPostgresSurvivorRepository(db *gorm.DB) *postgresSurvivorRepository {
	return &postgresSurvivorRepository{db}
}

func (repo *postgresSurvivorRepository) Save(surv *domain.Survivor) error {
	repo.db.Create(surv)
	tx := repo.db.Save(surv)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
