package domain

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	Username string
	Password string

	// Belongs to Survivor
	SurvivorID uint
	Survivor   *Survivor `gorm:"foreignKey:SurvivorID"`
}

type CredentialsRepository interface {
	Save(cred *Credentials) error
}
