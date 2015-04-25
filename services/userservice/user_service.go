package userservice

/*
CREATE TABLE users (
	id     serial primary key,
	email    varchar(255),
	password varchar(255)
);
*/

import (
	_ "database/sql"
	"encoding/json"
	"errors"
	_ "fmt"
	"github.com/davidamitchell/turbo-octo-avenger/context"
	"github.com/davidamitchell/turbo-octo-avenger/services"
	"github.com/davidamitchell/turbo-octo-avenger/util"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Email    string
	Password string
}

// Request for a User insert or update
type UserInsertParams struct {
	services.APIParams
	User
}

func NewUserInsertParams(requestID string) *UserInsertParams {
	return &UserInsertParams{
		APIParams: services.APIParams{requestID, ""},
		User:      User{}}
}

// Response from a Ping List operation
type UserInsertResponse struct {
	services.APIResponse
	User
}

func NewUserInsertResponse(requestID string) *UserInsertResponse {
	return &UserInsertResponse{APIResponse: services.NewAPIResponse(requestID)}
}

//
// Parse & Validate Insert Parameters required on every request
// - APIKey value, which must be a valid UUID
// - body containing User values
//
// Returns error or nil
func parseInsertParameters(r *http.Request, p *UserInsertParams) error {
	var err error
	if err = services.ParseAPIParameters(r, &p.APIParams); err != nil {
		return err
	}
	var body []byte
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(body, &p.User); err != nil {
		return err
	}
	if len(p.User.Email) == 0 {
		return errors.New("Required field 'email' is missing")
	}
	if len(p.User.Password) == 0 {
		return errors.New("Required field 'password' is missing")
	}
	return nil
}

//
// Create a handler to insert or update a  User resource
// - Decodes the URL params, Headers, Body
// - Calls the service
// - Encodes the response to Headers, Body, return code
//
func UserServiceInsertHandler() http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("severity=INFO RequestID=%v UserServiceInsertHandler", context.MustGetRequestId(r))
		var params = NewUserInsertParams(context.MustGetRequestId(r))
		var response = NewUserInsertResponse(context.MustGetRequestId(r))
		var err = parseInsertParameters(r, params)
		if err == nil {
			Insert(params, response)
		} else {
			// TODO create a function that adds error & sets HTTP status at same time
			response.Errors = append(response.Errors, services.Error{services.ErrorCodeBadRequest, err.Error()})
			response.HTTPStatus = http.StatusBadRequest
		}
		log.Printf("severity=INFO RequestID=%v UserServiceInsertHandler response='%s'", context.MustGetRequestId(r), util.MustMarshalJSON(response))
		w.WriteHeader(response.HTTPStatus)
		w.Write(util.MustMarshalJSON(response))
	}
	return f
}

func connect() *sqlx.DB {
	return sqlx.MustConnect("postgres", "postgres://root:@localhost/turbo_octo_avenger_development?sslmode=disable")
	// return sqlx.MustConnect("postgres", "postgres://root:root@localhost:5432/benchmarking"")
}

// Row in users table
type UserRow struct {
	ID       int
	Email    string
	Password string
}

func Insert(params *UserInsertParams, response *UserInsertResponse) {
	log.Printf("severity=DEBUG RequestID=%v User::Insert, email=%v password=%v", params.RequestID, params.User.Email, params.User.Password)
	db := connect()
	defer db.Close()
	var err error
	response.User = params.User
	err = db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		params.User.Email,
		params.User.Password).Scan(&response.User.ID)
	if err != nil {
		log.Printf("severity=ERROR RequestID=%v User::Insert query error: %v", params.RequestID, err)
		response.Errors = append(response.Errors, services.Error{services.ErrorCodeInternalServerError, err.Error()})
		response.HTTPStatus = http.StatusInternalServerError
		return
	}
	if response.User.ID == 0 {
		log.Printf("severity=ERROR RequestID=%v User::Insert New ID not returned! ", params.RequestID)
		response.Errors = append(response.Errors, services.Error{services.ErrorCodeInternalServerError, "Id not returned from insert"})
		response.HTTPStatus = http.StatusInternalServerError
		return
	}
	log.Printf("severity=INFO RequestID=%v User::Insert succeeded with id %v ", params.RequestID, response.User.ID)
}
