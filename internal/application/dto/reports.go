package dto

import "github.com/murilocosta/agartha/internal/domain"

type InfectedPercentageReportRead struct {
	InfectedPercentage float32 `json:"infected_percentage"`
}

type NonInfectedPercentageReportRead struct {
	NonInfectedPercentage float32 `json:"non_infected_percentage"`
}

type AverageResourcesPerSurvivorRead struct {
	Item        *ItemRead `json:"item"`
	ItemAverage float32   `json:"item_average"`
}

func ConvertToAverageResourcesPerSurvivorRead(report []*domain.AverageResourcesPerSurvivor) []*AverageResourcesPerSurvivorRead {
	var result []*AverageResourcesPerSurvivorRead
	for _, value := range report {
		result = append(result, &AverageResourcesPerSurvivorRead{
			Item:        ConvertToItemRead(value.Item),
			ItemAverage: value.ItemAverage,
		})
	}
	return result
}
