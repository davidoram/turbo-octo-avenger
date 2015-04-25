package services

import (
	"errors"
	_ "fmt"
	"github.com/davidoram/turbo-octo-avenger/context"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

// Generic Response
type APIResponse struct {
	RequestID  string // uuid as string
	HTTPStatus int
	Errors     []Error
}

type Error struct {
	Code    string
	Message string
}

func NewAPIResponse(requestID string) APIResponse {
	return APIResponse{
		RequestID:  requestID,
		HTTPStatus: http.StatusOK,
		Errors:     make([]Error, 0)}
}

const (
	ErrorCodeNotFound            = "not.found"             // 404 StatusNotFound
	ErrorCodeBadRequest          = "bad.request"           // 400 StatusBadRequest
	ErrorCodeInternalServerError = "internal.server.error" // 500 StatusInternalServerError
)

// Generic parameters passed to all services
type APIParams struct {
	RequestID string // uuid as string
	APIKey    string // uuid as string
	// Other parameter - perhaps an Access & Authorisation data?
}

// List Paramater validation & default constants
const (
	MinListOffset     = 0
	DefaultListOffset = 0

	MinListLimit     = 0
	MaxListLimit     = 200
	DefaultListLimit = 50
)

// Default parameters for a 'List' operation
type ListParams struct {
	APIParams
	Offset int
	Limit  int
}

//
// Parse & Validate the Generic Parameters required on every request
// - APIKey value, which must be a valid UUID (TODO and must a valid key)
//
// Returns error or nil
func ParseAPIParameters(r *http.Request, p *APIParams) error {
	log.Printf("severity=DEBUG RequestID=%v ParseAPIParameters %v", context.MustGetRequestId(r), r.Header)
	if apiKey, present := r.Header["X-Apikey"]; present {
		if key, err := uuid.ParseHex(apiKey[0]); err == nil {
			p.APIKey = key.String()
			return nil
		} else {
			return errors.New("Invalid X-Apikey header value")
		}
	} else {
		return errors.New("Missing X-Apikey header value")
	}
}

//
// Parse & Validate the Generic List Parameters required on every request
// - Offset : integer >= 0
// - Limit : integer between 1 and 200
//
// Returns error or nil
func ParseListParameters(r *http.Request, p *ListParams) error {
	e := ParseAPIParameters(r, &p.APIParams)
	if e != nil {
		return e
	}
	p.Offset = DefaultListOffset
	p.Limit = DefaultListLimit
	return nil
}
