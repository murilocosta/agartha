package domain

import "time"

type InfectionReport struct {
	Reportee   *Survivor
	Reported   *Survivor
	ReportedAt time.Time
}
