package core

type SetupErr error
type SystemErr error
type BusinessErr error

type ErrorMessage struct {
	ErrorType string
	Detail    string
	Status    uint
	Errors    []*ErrorDetail
}

type ErrorDetail struct {
	Name   string
	Reason string
}
