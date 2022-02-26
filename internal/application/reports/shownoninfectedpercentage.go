package reports

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type ShowNonInfectedPercentage struct {
	reportRepo domain.ReportRepository
}

func NewShowNonInfectedPercentage(reportRepo domain.ReportRepository) *ShowNonInfectedPercentage {
	return &ShowNonInfectedPercentage{reportRepo}
}

func (ucase *ShowNonInfectedPercentage) Invoke() (*dto.NonInfectedPercentageReportRead, error) {
	result, err := ucase.reportRepo.ShowNonInfectedPercentage()
	if err != nil {
		return nil, err
	}

	return &dto.NonInfectedPercentageReportRead{
		NonInfectedPercentage: result.Percentage,
	}, nil
}
