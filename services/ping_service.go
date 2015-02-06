package services

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// -----------------------------------------------------------------------------
//
// Ping Service
//

type PingService struct {
}

func (s *PingService) Name() string {
	return "ping"
}

func (s *PingService) Version() int {
	return 1
}

func (s *PingService) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "{ 'data': [1,2,3] }")
}

func (s *PingService) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "{ 'id': %v  }", p.ByName("id"))
}
