package services

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// -----------------------------------------------------------------------------
//
// Abstract interface to be implemented by all services
//

// Defines service type
type Service interface {
	Name() string
	Version() int

	List(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Show(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	//Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
