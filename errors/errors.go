package errors

import (
	"github.com/pkg/errors"
)

const (
	NotFound = iota + 1000
	InvalidParams
	UserNotFound
	InvalidPassword
	InternalServer
	PassWordNotAllowed
	NameExists
	NameNotExists

	Success     = 0
	Unknown     = -1
	GenericCode = 1
)

var (
	ErrUnknown         = newError(Unknown, "unknown error")
	ErrInvalidParams   = newError(InvalidParams, "invalid params")
	ErrUserNotFound    = newError(UserNotFound, "user not found")
	ErrInvalidPassword = newError(InvalidPassword, "invalid password")
	ErrInternalServer  = newError(InternalServer, "internal server error")
	ErrPassWord        = newError(PassWordNotAllowed, "password not allowed")
	ErrNameExists      = newError(NameExists, "the name Exists")
	ErrNameNotExists   = newError(NameNotExists, "the name not exists")
	ErrSuccess         = newError(Success, "success")
)

type GenericError struct {
	Code int
	Err  error
}

func (e GenericError) Error() string {
	return e.Err.Error()
}

func newError(code int, message string) GenericError {
	return GenericError{Code: code, Err: errors.New(message)}
}

func NewError(msg string) GenericError {
	return newError(GenericCode, msg)
}
