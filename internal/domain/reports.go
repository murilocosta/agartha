package domain

type InfectedPercentageReport struct {
	InfectedPercentage float32
}

type ReportRepository interface {
	ShowInfectedPercentage() (*InfectedPercentageReport, error)
}
