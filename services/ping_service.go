package services

import (
	_ "fmt"
	"github.com/davidoram/turbo-octo-avenger/config"
	"github.com/davidoram/turbo-octo-avenger/context"
	"github.com/davidoram/turbo-octo-avenger/util"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

//
// Create a handler to call the ping service
// - Decodes the URL params, Headers, Body
// - Calls the service
// - Encodes the response to Headers, Body, return code
//
func PingServiceListHandler() http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("severity=INFO RequestID=%v PingServiceListHandler", context.MustGetRequestId(r))
		var params = ListParams{}
		var response = NewPingResponse(context.MustGetRequestId(r))
		var err = ParseListParameters(r, &params)
		if err == nil {
			List(&params, response)
		} else {
			response.Errors = append(response.Errors, Error{ErrorCodeBadRequest, err.Error()})
			response.HTTPStatus = http.StatusBadRequest
		}
		w.WriteHeader(response.HTTPStatus)
		w.Write(util.MustMarshalJSON(response))

	}
	return f
}

func connect() *sqlx.DB {
	dbUrl, _ := config.DataBaseURL("development")
	return sqlx.MustConnect("postgres", dbUrl) //TODO should be read from somewhere
}

// Response from a Ping List operation
type PingResponse struct {
	APIResponse
	Message string
}

func NewPingResponse(requestID string) *PingResponse {
	return &PingResponse{
		APIResponse: NewAPIResponse(requestID),
		Message:     ""}
}

// Row in ping table
type PingRow struct {
	Message string
}

func List(params *ListParams, response *PingResponse) {

	db := connect()
	defer db.Close()
	log.Printf("severity=DEBUG RequestID=%v Ping::List. db=%v", params.RequestID, db)
	var row PingRow
	err := db.QueryRowx("SELECT message FROM ping LIMIT 1").StructScan(&row)
	if err != nil {
		log.Printf("severity=ERROR RequestID=%v Ping::List query error: %v", params.RequestID, err)
		response.Errors = append(response.Errors, Error{ErrorCodeInternalServerError, err.Error()})
		response.HTTPStatus = http.StatusInternalServerError
		return
	}
	log.Printf("severity=DEBUG RequestID=%v Ping::List query ok", params.RequestID)
	response.Message = row.Message
}
