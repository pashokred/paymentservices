package domain

import (
	"encoding/json"
	"net/http"
)

type ButtonErrorInterface interface {
	Status() int
	Message() string
}
type ButtonError struct {
	Code      		 int           `json:"code"`
	ErrorMessage     string        `json:"error"`
}

func (b *ButtonError) Status() int {
	return b.Code
}
func (b *ButtonError) Message() string {
	return b.ErrorMessage
}

func NewButtonError(statusCode int, message string) ButtonErrorInterface {
	return &ButtonError{
		Code:         statusCode,
		ErrorMessage: message,
	}
}
func NewBadRequestError(message string) ButtonErrorInterface {
	return &ButtonError{
		Code: http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewForbiddenError(message string) ButtonErrorInterface {
	return &ButtonError{
		Code: http.StatusForbidden,
		ErrorMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (ButtonErrorInterface, error) {
	var result ButtonError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}


