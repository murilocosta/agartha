package persistence

import (
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
	sqlQuery := `
		SELECT 
			CAST((SELECT COALESCE(COUNT(s1.id), 0) FROM survivors AS s1 WHERE s1.infected IS TRUE) AS DECIMAL)
			/
			(SELECT COALESCE(COUNT(s2.id), 1) FROM survivors AS s2) 
		AS infected_percentage
	`

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
	sqlQuery := `
		SELECT 
			CAST((SELECT COALESCE(COUNT(s1.id), 0) FROM survivors AS s1 WHERE s1.infected IS FALSE) AS DECIMAL)
			/
			(SELECT COALESCE(COUNT(s2.id), 1) FROM survivors AS s2) 
		AS infected_percentage
	`

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

func (repo *postgresReportRepository) ShowAverageResourcesPerSurvivor() ([]*domain.AverageResourcesPerSurvivor, error) {
	sqlQuery := `
		SELECT it.name, it.icon, it.price, it.rarity, AVG(res.quantity) AS item_average 
		FROM survivors AS surv
		INNER JOIN inventories AS inv ON inv.owner_id = surv.id
		INNER JOIN resources AS res ON res.inventory_id = inv.id
		INNER JOIN items AS it ON it.id = res.item_id 
		GROUP BY 1, 2, 3, 4
		ORDER BY it.price DESC
	`

	rows, err := repo.db.Raw(sqlQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.AverageResourcesPerSurvivor

	var itemAverage float32
	for rows.Next() {
		item := domain.Item{}

		err = rows.Scan(&item.Name, &item.Icon, &item.Price, &item.Rarity, &itemAverage)
		if err != nil {
			return nil, err
		}

		resultItem := &domain.AverageResourcesPerSurvivor{
			Item:        &item,
			ItemAverage: itemAverage,
		}
		result = append(result, resultItem)
	}

	return result, nil
}
