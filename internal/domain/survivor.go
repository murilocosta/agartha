package domain

import "gorm.io/gorm"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)

type Location struct {
	Longitude float32 `json:"longitude" validate:"required,longitude"`
	Latitude  float32 `json:"latitude" validate:"required,latitude"`
	Timezone  string  `json:"timezone"`
}

type Survivor struct {
	gorm.Model
	Name         string
	Gender       Gender
	LastLocation *Location `gorm:"embedded;embeddedPrefix:location_"`
	Infected     bool
	Deceased     bool

	// Has one Credentials
	Credentials *Credentials

	// Has one Inventory
	Inventory *Inventory `gorm:"foreignKey:OwnerID"`
}

type SurvivorRepository interface {
	Save(surv *Survivor) error
	UpdateLastLocation(surv *Survivor, loc *Location) error

	FindByID(survID uint) (*Survivor, error)
}
