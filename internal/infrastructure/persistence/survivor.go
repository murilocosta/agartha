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
