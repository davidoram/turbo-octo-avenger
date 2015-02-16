package ipc

import (
	"github.com/jmoiron/sqlx"
	"github.com/nu7hatch/gouuid"
)

// -----------------------------------------------------------------------------
//
// Abstract interface to be implemented by all services
//
type ServiceMethod int

const (
	ListMethod   ServiceMethod = 0
	ShowMethod   ServiceMethod = 1
	CreateMethod ServiceMethod = 2
	UpdateMethod ServiceMethod = 3
	DeleteMethod ServiceMethod = 4
)

// Defines service type
type Service interface {
	Name() string
	Version() int

	List(requestId *uuid.UUID, db *sqlx.DB, rv *ReturnValue) error
	// List(w http.ResponseWriter, r *http.Request)
	// Show(w http.ResponseWriter, r *http.Request)
	//Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
