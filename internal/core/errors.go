package core

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func NewErrorMessage(errType string, detail string, status uint) *ErrorMessage {
	return &ErrorMessage{
		ErrorType: errType,
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
