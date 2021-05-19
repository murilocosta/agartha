package dto

import "github.com/murilocosta/agartha/internal/core"

// Validation error codes
const (
	RegisterSurvivorFailed   core.ErrorTypeCode = "AGV-001"
	UpdateLastLocationFailed core.ErrorTypeCode = "AGV-002"
)

// Business rule error codes
const (
	ItemNotFound     core.ErrorTypeCode = "AGB-001"
	SurvivorNotFound core.ErrorTypeCode = "AGB-002"
)
