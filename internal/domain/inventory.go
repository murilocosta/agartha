package domain

type Inventory struct {
	Owner     *Survivor
	Resources []*Resource
}
