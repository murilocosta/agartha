package reports

import (
	"github.com/murilocosta/agartha/internal/application/dto"
	"github.com/murilocosta/agartha/internal/domain"
)

type ShowTotalResourceQuantityLost struct {
	reportRepo domain.ReportRepository
}

func NewShowTotalResourceQuantityLost(reportRepo domain.ReportRepository) *ShowTotalResourceQuantityLost {
	return &ShowTotalResourceQuantityLost{reportRepo}
}

func (ucase *ShowTotalResourceQuantityLost) Invoke() ([]*dto.ShowTotalResourceQuantityLostRead, error) {
	result, err := ucase.reportRepo.ShowTotalResourceQuantityLost()
	if err != nil {
		return nil, err
	}

	report := dto.ConvertToShowTotalResourceQuantityLostRead(result)
	return report, nil
}
