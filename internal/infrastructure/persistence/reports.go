package persistence

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/domain"
)

type postgresReportRepository struct {
	db *gorm.DB
}

func NewPostgresReportRepository(db *gorm.DB) domain.ReportRepository {
	return &postgresReportRepository{db}
}

func (repo *postgresReportRepository) ShowInfectedPercentage() (*domain.InfectedReport, error) {
	sqlPart1 := "SELECT COALESCE(COUNT(s1.id), 0) FROM survivors AS s1 WHERE s1.infected IS TRUE"
	sqlPart2 := "SELECT COALESCE(COUNT(s2.id), 1) FROM survivors AS s2"
	sqlQuery := fmt.Sprintf("SELECT CAST((%s) AS DECIMAL) / (%s) AS infected_percentage", sqlPart1, sqlPart2)

	rows, err := repo.db.Raw(sqlQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var infectedPercentage float32
	for rows.Next() {
		err = rows.Scan(&infectedPercentage)
		if err != nil {
			return nil, err
		}
	}

	return &domain.InfectedReport{Percentage: infectedPercentage}, nil
}

func (repo *postgresReportRepository) ShowNonInfectedPercentage() (*domain.InfectedReport, error) {
	sqlPart1 := "SELECT COALESCE(COUNT(s1.id), 0) FROM survivors AS s1 WHERE s1.infected IS FALSE"
	sqlPart2 := "SELECT COALESCE(COUNT(s2.id), 1) FROM survivors AS s2"
	sqlQuery := fmt.Sprintf("SELECT CAST((%s) AS DECIMAL) / (%s) AS infected_percentage", sqlPart1, sqlPart2)

	rows, err := repo.db.Raw(sqlQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nonInfectedPercentage float32
	for rows.Next() {
		err = rows.Scan(&nonInfectedPercentage)
		if err != nil {
			return nil, err
		}
	}

	return &domain.InfectedReport{Percentage: nonInfectedPercentage}, nil
}
