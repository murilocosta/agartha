package persistence

import (
	"github.com/murilocosta/agartha/internal/domain"
	"gorm.io/gorm"
)

type postgresCredentialsRepository struct {
	db *gorm.DB
}

func NewPostgresCredentialsRepository(db *gorm.DB) domain.CredentialsRepository {
	return &postgresCredentialsRepository{db}
}

func (repo *postgresCredentialsRepository) Save(cred *domain.Credentials) error {
	tx := repo.db.Save(cred)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *postgresCredentialsRepository) FindByUsername(username string) (*domain.Credentials, error) {
	var cred domain.Credentials

	err := repo.db.Where("username = ?", username).First(&cred).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&cred).Association("Survivor").Find(&cred.Survivor)
	if err != nil {
		return nil, err
	}

	return &cred, nil
}
