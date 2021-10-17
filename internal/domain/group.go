package domain

import "gorm.io/gorm"

type GroupMember struct {
	gorm.Model
	Role string

	// Has one Credentials
	MemberID uint
	Member   *Credentials `gorm:"foreignKey:MemberID"`

	// Belongs to Group
	GroupID uint
	Group   *Group
}

type Group struct {
	gorm.Model
	Name string

	Members []*GroupMember
}
