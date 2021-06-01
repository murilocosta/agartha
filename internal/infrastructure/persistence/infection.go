package persistence

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresInfectionRepository struct {
	db *gorm.DB
}

func NewPostgresInfectionRepository(db *gorm.DB) *postgresInfectionRepository {
	return &postgresInfectionRepository{db}
}

func (repo *postgresInfectionRepository) Save(reportee *domain.Survivor, reported *domain.Survivor, annotation string) error {
	infection := &domain.InfectionReport{
		Reportee:   reportee,
		Reported:   reported,
		Annotation: annotation,
	}

	err := repo.db.Save(infection).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *postgresInfectionRepository) CountReports(reportedID uint) (int64, error) {
	var count int64
	tx := repo.db.Model(&domain.InfectionReport{}).Where("reported_id = ?", reportedID).Count(&count)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

func (repo *postgresInfectionRepository) HasReported(reporteeID uint, reportedID uint) bool {
	var found domain.InfectionReport

	tx := repo.db.Find(&found, &domain.InfectionReport{ReporteeID: reporteeID, ReportedID: reportedID})
	if tx.Error != nil {
		return false
	}

	if found.ID == 0 {
		return false
	}

	return true
}
