package dto

import "github.com/murilocosta/agartha/internal/domain"

type InfectedPercentageReportRead struct {
	InfectedPercentage float32 `json:"infected_percentage"`
}

type NonInfectedPercentageReportRead struct {
	NonInfectedPercentage float32 `json:"non_infected_percentage"`
}

type AverageResourcesPerSurvivorRead struct {
	ItemName    string            `json:"item_name"`
	ItemIcon    string            `json:"item_icon"`
	ItemPrice   int32             `json:"item_price"`
	ItemRarity  domain.ItemRarity `json:"item_rarity"`
	ItemAverage float32           `json:"item_average"`
}

func ConvertToAverageResourcesPerSurvivorRead(report []*domain.AverageResourcesPerSurvivor) []*AverageResourcesPerSurvivorRead {
	var result []*AverageResourcesPerSurvivorRead
	for _, value := range report {
		result = append(result, &AverageResourcesPerSurvivorRead{
			ItemName:    value.Item.Name,
			ItemIcon:    value.Item.Icon,
			ItemPrice:   value.Item.Price,
			ItemRarity:  value.Item.Rarity,
			ItemAverage: value.ItemAverage,
		})
	}
	return result
}
