package domain

type InfectedReport struct {
	Percentage float32
}

type AverageResourcesPerSurvivor struct {
	Item        *Item
	ItemAverage float32
}

type ReportRepository interface {
	ShowInfectedPercentage() (*InfectedReport, error)
	ShowNonInfectedPercentage() (*InfectedReport, error)
	ShowAverageResourcesPerSurvivor() ([]*AverageResourcesPerSurvivor, error)
}
