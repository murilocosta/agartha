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

type ShowTotalResourceQuantityLostRead struct {
	Item             *ItemRead `json:"item"`
	ItemQuantityLost int32     `json:"item_quantity_lost"`
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

func ConvertToShowTotalResourceQuantityLostRead(report []*domain.TotalResourceQuantityLost) []*ShowTotalResourceQuantityLostRead {
	var result []*ShowTotalResourceQuantityLostRead
	for _, value := range report {
		result = append(result, &ShowTotalResourceQuantityLostRead{
			Item:             ConvertToItemRead(value.Item),
			ItemQuantityLost: value.ItemQuantityLost,
		})
	}
	return result
}
