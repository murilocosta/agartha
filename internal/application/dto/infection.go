package dto

type ReportedInfection struct {
	ReporteeID uint   `json:"reportee_id" validate:"required"`
	ReportedID uint   `json:"reported_id" validate:"required"`
	Annotation string `json:"annotation"`
}
