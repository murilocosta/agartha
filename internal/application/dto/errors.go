package dto

import "github.com/murilocosta/agartha/internal/core"

// Validation error codes
const (
	RegisterSurvivorFailed       core.ErrorTypeCode = "AGV-001"
	UpdateLastLocationFailed     core.ErrorTypeCode = "AGV-002"
	ReportInfectedSurvivorFailed core.ErrorTypeCode = "AGV-003"
)

// Business rule error codes
const (
	ItemNotFound                    core.ErrorTypeCode = "AGB-001"
	SurvivorNotFound                core.ErrorTypeCode = "AGB-002"
	SurvivorCannotBeFlagged         core.ErrorTypeCode = "AGB-003"
	SurvivorAlreadyFlagged          core.ErrorTypeCode = "AGB-004"
	SurvivorInventoryTransferFailed core.ErrorTypeCode = "AGB-005"
	SurvivorInventoryPriceExceeded  core.ErrorTypeCode = "AGB-011"
	TradeResourceNotFound           core.ErrorTypeCode = "AGB-006"
	TradeResourcePriceNotEquivalent core.ErrorTypeCode = "AGB-007"
	TradeStatusIsInvalid            core.ErrorTypeCode = "AGB-008"
	TradeResourceQuantityNotEnough  core.ErrorTypeCode = "AGB-009"
	TradeCanOnlyBeCancelledBySender core.ErrorTypeCode = "AGB-010"
)

func ErrorDetailBuilder(field string, errorType string) *core.ErrorDetail {
	return &core.ErrorDetail{
		Name:   field,
		Reason: errorType,
	}
}
