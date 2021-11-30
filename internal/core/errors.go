package core

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SetupErr error
type SystemErr error
type BusinessErr error
type ErrorBuilder func(string, string) *ErrorDetail

type ErrorMessage struct {
	ErrorType string         `json:"errorType"`
	Detail    string         `json:"detail"`
	Status    uint           `json:"status"`
	Errors    []*ErrorDetail `json:"errors,omitempty"`
}

type ErrorDetail struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ErrorTypeCode string

func NewErrorMessage(errType ErrorTypeCode, detail string, status uint) *ErrorMessage {
	return &ErrorMessage{
		ErrorType: string(errType),
		Detail:    detail,
		Status:    status,
	}
}

func (m *ErrorMessage) AddErrorDetail(err error, builder ErrorBuilder) {
	for _, err := range err.(validator.ValidationErrors) {
		detail := builder(err.StructField(), err.ActualTag())
		m.Errors = append(m.Errors, detail)
	}
}

func (m *ErrorMessage) Error() string {
	errJson, _ := json.Marshal(m)
	return string(errJson)
}

func GetErrorMessage(err error) *ErrorMessage {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return GetDatabaseError(err)
	}

	switch err.(type) {
	case *ErrorMessage:
		return GetValidationError(err)
	default:
		return GetSystemError(err)
	}
}

func GetValidationError(err error) *ErrorMessage {
	var msg ErrorMessage
	jsonMsg := err.Error()
	json.Unmarshal([]byte(jsonMsg), &msg)
	return &msg
}

func GetSystemError(err error) *ErrorMessage {
	return &ErrorMessage{
		ErrorType: "AGS-000",
		Detail:    err.Error(),
		Status:    http.StatusInternalServerError,
	}
}

func GetDatabaseError(err error) *ErrorMessage {
	return &ErrorMessage{
		ErrorType: "AGS-001",
		Detail:    "Access to this resource is denied",
		Status:    http.StatusForbidden,
	}
}
