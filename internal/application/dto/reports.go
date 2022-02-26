package dto

type InfectedPercentageReportRead struct {
	InfectedPercentage float32 `json:"infected_percentage"`
}

type NonInfectedPercentageReportRead struct {
	NonInfectedPercentage float32 `json:"non_infected_percentage"`
}
