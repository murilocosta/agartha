package domain

import (
	"gorm.io/gorm"
)

const (
	FlagInfectedSurvivorMax int64 = 5
)

type InfectionReport struct {
	gorm.Model
	Annotation string

	// Has one reportee
	ReporteeID uint
	Reportee   *Survivor `gorm:"foreignKey:ReporteeID"`

	// Has one reported
	ReportedID uint
	Reported   *Survivor `gorm:"foreignKey:ReportedID"`
}

type InfectionRepository interface {
	Save(reportee *Survivor, reported *Survivor, Annotation string) error

	CountReports(reportedID uint) (int64, error)
	HasReported(reporteeID uint, reportedID uint) bool
}
