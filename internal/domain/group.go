package domain

type GroupMember struct {
	Role   string
	Member *Survivor
}

type Group struct {
	Name    string
	Members []*GroupMember
}
