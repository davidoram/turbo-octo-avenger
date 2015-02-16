package ipc

import (
	"github.com/nu7hatch/gouuid"
	// "net/http"
)

type ErrorSource string

const (
	MiddlewareError ErrorSource = "middleware"
	ServiceError    ErrorSource = "service"
)

type ReturnValue struct {
	StatusCode int
	RequestID  *uuid.UUID
	Errors     []*error
	Data       []interface{}
}
