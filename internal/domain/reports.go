package domain

type InfectedReport struct {
	Percentage float32
}

type AverageResourcesPerSurvivor struct {
	Item        *Item
	ItemAverage float32
}

type TotalResourceQuantityLost struct {
	Item             *Item
	ItemQuantityLost int32
}

type ReportRepository interface {
	ShowInfectedPercentage() (*InfectedReport, error)
	ShowNonInfectedPercentage() (*InfectedReport, error)
	ShowAverageResourcesPerSurvivor() ([]*AverageResourcesPerSurvivor, error)
	ShowTotalResourceQuantityLost() ([]*TotalResourceQuantityLost, error)
}
