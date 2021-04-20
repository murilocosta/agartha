package domain

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)

type Location struct {
	Longitude float64
	Latitude  float64
	Timezone  string
}

type Survivor struct {
	Credentials  *Credential
	Name         string
	Gender       Gender
	LastLocation *Location `gorm:"embedded;embeddedPrefix:location_"`
	Infected     bool
	Deceased     bool
}
