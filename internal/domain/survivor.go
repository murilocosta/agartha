package domain

import (
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/core"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)

type SurvivorFilter struct {
	Name      string                `form:"name"`
	Sort      core.DatabaseSortType `form:"sort"`
	Page      int                   `form:"page"`
	PageItems int                   `form:"page_items"`
}

func NewSurvivorFilter(name string, sort string, page int) *SurvivorFilter {
	return &SurvivorFilter{
		Name:      name,
		Sort:      core.DatabaseSortType(sort),
		Page:      page,
		PageItems: 15,
	}
}

type Location struct {
	Longitude float32 `json:"longitude" validate:"required,longitude"`
	Latitude  float32 `json:"latitude" validate:"required,latitude"`
	Timezone  string  `json:"timezone"`
}

type Survivor struct {
	gorm.Model
	Name         string
	Gender       Gender
	Age          uint
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
	FindAll(filter *SurvivorFilter) ([]*Survivor, error)
}
