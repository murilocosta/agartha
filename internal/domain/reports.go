package domain

type InfectedReport struct {
	Percentage float32
}

type ReportRepository interface {
	ShowInfectedPercentage() (*InfectedReport, error)
	ShowNonInfectedPercentage() (*InfectedReport, error)
}
