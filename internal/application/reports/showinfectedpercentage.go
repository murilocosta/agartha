package reports

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type ShowInfectedPercentage struct {
	reportRepo domain.ReportRepository
}

func NewShowInfectedPercentage(reportRepo domain.ReportRepository) *ShowInfectedPercentage {
	return &ShowInfectedPercentage{reportRepo}
}

func (ucase *ShowInfectedPercentage) Invoke() (*dto.InfectedPercentageReportRead, error) {
	result, err := ucase.reportRepo.ShowInfectedPercentage()
	if err != nil {
		return nil, err
	}

	return &dto.InfectedPercentageReportRead{
		InfectedPercentage: result.Percentage,
	}, nil
}
