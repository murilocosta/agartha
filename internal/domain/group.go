package domain

type GroupMember struct {
	Role   string
	Member *Credentials
}

type Group struct {
	Name    string
	Members []*GroupMember
}
