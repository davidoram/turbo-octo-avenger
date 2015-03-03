package services

import (
	"errors"
	_ "fmt"
	"github.com/davidoram/turbo-octo-avenger/context"
	"github.com/davidoram/turbo-octo-avenger/util"
	"github.com/jmoiron/sqlx"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

// Generic Response
type APIResponse struct {
	RequestID  *uuid.UUID
	HTTPStatus int
	Errors     []error
}

// Generic parameters passed to all services
type APIParams struct {
	RequestID *uuid.UUID
	APIKey    *uuid.UUID
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
	log.Printf("RequestID=%v ParseAPIParameters %v", context.MustGetRequestId(r), r.Header)
	if apiKey, present := r.Header["X-Apikey"]; present {
		var err error
		if p.APIKey, err = uuid.ParseHex(apiKey[0]); err == nil {
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

//
// Create a handler to call the ping service
// - Decodes the URL params, Headers, Body
// - Calls the service
// - Encodes the response to Headers, Body, return code
//
func PingServiceListHandler() http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("RequestID=%v PingServiceListHandler", context.MustGetRequestId(r))
		var params = ListParams{}
		var response = NewPingResponse(context.MustGetRequestId(r))
		var err = ParseListParameters(r, &params)
		if err == nil {
			List(&params, response)
		} else {
			response.Errors = append(response.Errors, err)
			response.HTTPStatus = http.StatusBadRequest
		}
		w.WriteHeader(response.HTTPStatus)
		w.Write(util.MustMarshalJSON(response))

	}
	return f
}

func connect() *sqlx.DB {
	return sqlx.MustConnect("postgres", "postgres://davidoram:@localhost/turbo-octo-avenger-development?sslmode=disable")
}

// Response from a Ping List operation
type PingResponse struct {
	APIResponse
	Message string
}

func NewPingResponse(requestID *uuid.UUID) *PingResponse {
	return &PingResponse{
		APIResponse: APIResponse{
			RequestID:  requestID,
			HTTPStatus: http.StatusOK,
			Errors:     make([]error, 0)},
		Message: ""}
}

// Row in ping table
type PingRow struct {
	Message string
}

func List(params *ListParams, response *PingResponse) {

	db := connect()
	log.Printf("RequestID=%v Ping::List. db=%v", params.RequestID, db)
	var row PingRow
	err := db.QueryRowx("SELECT message FROM ping LIMIT 1").StructScan(&row)
	if err != nil {
		log.Printf("RequestID=%v Ping::List query error: %v", params.RequestID, err)
		response.Errors = append(response.Errors, err)
		response.HTTPStatus = http.StatusInternalServerError
		return
	}
	log.Printf("RequestID=%v Ping::List query ok", params.RequestID)
	response.Message = row.Message
}
