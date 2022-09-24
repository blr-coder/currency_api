package postgres

import (
	"encoding/json"
	"errors"
	"reflect"
)

type ErrorCode uint

const (
	ErrorCodeOK ErrorCode = iota
	ErrorCodeCanceled
	ErrorCodeUnknown
	ErrorCodeInvalidArgument
	ErrorCodeDeadlineExceeded
	ErrorCodeNotFound
	ErrorCodeAlreadyExists
	ErrorCodePermissionDenied
	ErrorCodeResourceExhausted
	ErrorCodeFailedPrecondition
	ErrorCodeAborted
	ErrorCodeOutOfRange
	ErrorCodeUnimplemented
	ErrorCodeInternal
	ErrorCodeUnavailable
	ErrorCodeDataLoss
	ErrorCodeUnauthenticated
)

type Error struct {
	Code    ErrorCode         `json:"code"`
	Message string            `json:"message"`
	Params  map[string]string `json:"params"`
}

func (e Error) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (e *Error) Is(tgt error) bool {
	var target *Error
	if !errors.As(tgt, &target) {
		return false
	}
	return reflect.DeepEqual(e, target)
}

func (e *Error) SetCode(code ErrorCode) {
	e.Code = code
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetParams(params map[string]string) {
	e.Params = params
}

func (e *Error) AddParam(key string, value string) {
	e.Params[key] = value
}

func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Params:  map[string]string{},
	}
}

func NewUnexpectedBehaviorError(details string) *Error {
	return &Error{
		Code:    ErrorCodeInternal,
		Message: "Unexpected behavior.",
		Params: map[string]string{
			"details": details,
		},
	}
}

func NewInvalidFormError() *Error {
	return NewError(ErrorCodeInvalidArgument, "The form sent is not valid, please correct the errors below.")
}

func NewCurrencyPairNotFound() *Error {
	e := NewError(ErrorCodeNotFound, "CurrencyPair not found.")
	return e
}

func NewCurrencyPairAlreadyExist() *Error {
	e := NewError(ErrorCodeFailedPrecondition, "CurrencyPair already exist.")
	return e
}
