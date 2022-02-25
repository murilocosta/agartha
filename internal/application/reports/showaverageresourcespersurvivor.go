package reports

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type ShowAverageResourcesPerSurvivor struct {
	reportRepo domain.ReportRepository
}

func NewShowAverageResourcesPerSurvivor(reportRepo domain.ReportRepository) *ShowAverageResourcesPerSurvivor {
	return &ShowAverageResourcesPerSurvivor{reportRepo}
}

func (ucase *ShowAverageResourcesPerSurvivor) Invoke() ([]*dto.AverageResourcesPerSurvivorRead, error) {
	result, err := ucase.reportRepo.ShowAverageResourcesPerSurvivor()
	if err != nil {
		return nil, err
	}

	report := dto.ConvertToAverageResourcesPerSurvivorRead(result)
	return report, nil
}
